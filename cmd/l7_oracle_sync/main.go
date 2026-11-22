// PROJECT NAME: ThitNueaHub-Core-Master // FOLDER PATH: /cmd/l7_oracle_sync/ // FILE NAME: main.go
package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ARTHITSIANGWAN/ThitNueaHub-Core-Master/backend/core"
)

func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	fmt.Println("[L7 Oracle] Starting autonomous 5S audit, verification, and Git sync...")

	forbidden := []string{
		"tnh-gripen-" + "sovereign-secret",
		"L1-" + "COMMANDER",
		"v" + "83",
		"0." + "xxms",
	}

	var violations []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == ".git" || path == filepath.Join("cmd", "l7_oracle_sync") {
				return filepath.SkipDir
			}
			return nil
		}
		if path != "go.mod" && d.Name() == "go.mod" {
			violations = append(violations, "nested go.mod detected: "+path)
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil
		}
		content := string(data)
		for _, bad := range forbidden {
			if strings.Contains(content, bad) {
				violations = append(violations, fmt.Sprintf("%s contains forbidden token %q", path, bad))
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("[L7 Oracle] Walk error: %v\n", err)
		os.Exit(1)
	}
	if len(violations) > 0 {
		for _, v := range violations {
			fmt.Printf("[L7 Oracle] VIOLATION: %s\n", v)
		}
		os.Exit(1)
	}
	fmt.Println("[L7 Oracle] 5S Context & Structure Scan: CLEAN")

	probe := []byte("L7-ORACLE|HEALTH_CHECK")
	sig, err := core.SignPayload(probe)
	if err != nil {
		fmt.Printf("[L7 Oracle] L8 Glock sign failed: %v\n", err)
		os.Exit(1)
	}
	ok, err := core.VerifyPayload(probe, sig)
	if err != nil || !ok {
		fmt.Printf("[L7 Oracle] L8 Glock verify failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[L7 Oracle] L8 Glock Perimeter Check (bind=%s): PASSED\n", core.GetBindAddr())

	if err := runCmd("go", "vet", "./..."); err != nil {
		fmt.Printf("[L7 Oracle] go vet failed: %v\n", err)
		os.Exit(1)
	}
	if err := runCmd("go", "test", "./..."); err != nil {
		fmt.Printf("[L7 Oracle] go test failed: %v\n", err)
		os.Exit(1)
	}

	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		fmt.Printf("[L7 Oracle] git status failed: %v\n", err)
		os.Exit(1)
	}
	if len(bytes.TrimSpace(out)) == 0 {
		fmt.Println("[L7 Oracle] Working tree clean. Nothing to push.")
		return
	}

	if err := runCmd("git", "add", "-A"); err != nil {
		os.Exit(1)
	}
	if err := runCmd("git", "commit", "-m", "feat(core): deploy L7 Oracle sync runner and L8 env-based HMAC entrypoints"); err != nil {
		os.Exit(1)
	}
	if err := runCmd("git", "push", "origin", "main"); err != nil {
		fmt.Printf("[L7 Oracle] git push failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("[L7 Oracle] MISSION COMPLETE: Verified & pushed to GitHub main.")
}
