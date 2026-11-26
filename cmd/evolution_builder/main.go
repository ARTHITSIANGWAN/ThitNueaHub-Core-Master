// PROJECT NAME: ThitNueaHub-Core-Master (tnh_ai)
// FOLDER PATH: /cmd/evolution_builder/
// FILE NAME: main.go
// LICENSE: MIT License © 2026 ThitNueaHub (Arthit Siangwan)

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	validLayerRegex = regexp.MustCompile(`^L(0[1-9]|1[0-1])$`)
	stageFiles      = []string{"age05.png", "age15.png", "age24.png", "age45.png", "age65.png"}
)

func main() {
	baseDir := flag.String("base", "./assets/agents", "Base directory containing L01-L11 folders")
	layerID := flag.String("layer", "", "Target agent layer ID (L01 to L11)")
	outFile := flag.String("out", "", "Output filename inside the layer folder (default: <LAYER>_evolution_sheet.png)")
	flag.Parse()

	target := strings.ToUpper(strings.TrimSpace(*layerID))
	if !validLayerRegex.MatchString(target) {
		fmt.Fprintf(os.Stderr, "[SECURITY BLOCK] Invalid layer ID %q: must be L01 through L11\n", target)
		os.Exit(1)
	}

	cleanBase, err := filepath.Abs(filepath.Clean(*baseDir))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Invalid base path: %v\n", err)
		os.Exit(1)
	}

	layerDir, err := findLayerFolder(cleanBase, target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}

	var args []string
	for _, stage := range stageFiles {
		imgPath := filepath.Join(layerDir, stage)
		if !isWithinDir(layerDir, imgPath) {
			fmt.Fprintf(os.Stderr, "[SECURITY BLOCK] Path traversal detected on %s\n", stage)
			os.Exit(1)
		}
		if _, err := os.Stat(imgPath); err != nil {
			fmt.Fprintf(os.Stderr, "[MISSING ASSET] Required stage image not found: %s\n", imgPath)
			os.Exit(1)
		}
		args = append(args, "-i", imgPath)
	}

	outputName := strings.TrimSpace(*outFile)
	if outputName == "" {
		outputName = fmt.Sprintf("%s_evolution_sheet.png", target)
	}
	outputPath := filepath.Join(layerDir, filepath.Base(outputName))
	if !isWithinDir(layerDir, outputPath) {
		fmt.Fprintf(os.Stderr, "[SECURITY BLOCK] Invalid output path\n")
		os.Exit(1)
	}

	filterComplex := "[0:v]scale=400:800:force_original_aspect_ratio=increase,crop=400:800[v0];" +
		"[1:v]scale=400:800:force_original_aspect_ratio=increase,crop=400:800[v1];" +
		"[2:v]scale=400:800:force_original_aspect_ratio=increase,crop=400:800[v2];" +
		"[3:v]scale=400:800:force_original_aspect_ratio=increase,crop=400:800[v3];" +
		"[4:v]scale=400:800:force_original_aspect_ratio=increase,crop=400:800[v4];" +
		"[v0][v1][v2][v3][v4]hstack=inputs=5[out]"

	args = append(args,
		"-filter_complex", filterComplex,
		"-map", "[out]",
		"-frames:v", "1",
		"-update", "1",
		"-map_metadata", "-1",
		"-y", outputPath,
	)

	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("[TNH-EVOLUTION] Building 5-stage clean sheet for %s...\n", target)
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "[FFMPEG ERROR] %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[SUCCESS] Clean metadata-stripped sheet saved to: %s\n", outputPath)
}

func findLayerFolder(baseDir, layerPrefix string) (string, error) {
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return "", fmt.Errorf("cannot read base directory %s: %w", baseDir, err)
	}
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(strings.ToUpper(entry.Name()), layerPrefix+"_") {
			candidate := filepath.Join(baseDir, entry.Name())
			if isWithinDir(baseDir, candidate) {
				return candidate, nil
			}
		}
	}
	return "", fmt.Errorf("folder starting with %s_ not found in %s", layerPrefix, baseDir)
}

func isWithinDir(parent, target string) bool {
	rel, err := filepath.Rel(parent, target)
	if err != nil {
		return false
	}
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(os.PathSeparator))
}
