// cmd/orchestra/main.go
package main

import (
	"fmt"

	"thitnueahub/internal/core"
)

func main() {
	fmt.Println("🌟 TNH AI Orchestra v8.3 - 9th Era of Fire")
	fmt.Println("Private Fortress | Zero-Garbage Go Engine")
	fmt.Println("========================================")

	orchestrator := core.NewOrchestrator()

	// Register Guardians
	orchestrator.RegisterAll()

	orchestrator.Start()
}
