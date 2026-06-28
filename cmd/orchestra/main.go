package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ==================== State ====================

type ThitnueahubState struct {
	Messages        []string
	OverallProgress float64
	OverallStatus   string
	ActiveIssues    []string
	DebateHistory   []string
	GuardianStatus  map[string]string
	Metadata        map[string]interface{}
}

// Guardian Interface
type Guardian interface {
	Name() string
	Execute(ctx context.Context, state *ThitnueahubState) error
}

// ==================== Commander ====================

type Commander struct {
	guardians []Guardian
	state     *ThitnueahubState
}

func NewCommander() *Commander {
	return &Commander{
		guardians: make([]Guardian, 0),
		state: &ThitnueahubState{
			GuardianStatus: make(map[string]string),
			Metadata:       make(map[string]interface{}),
		},
	}
}

func (c *Commander) Register(g Guardian) {
	c.guardians = append(c.guardians, g)
	c.state.GuardianStatus[g.Name()] = "registered"
	fmt.Printf("✅ Registered Guardian: %s\n", g.Name())
}

func (c *Commander) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("\n🛑 Shutting down gracefully...")
		cancel()
	}()

	fmt.Println("🚀 TNH AI Orchestra Started - L1 Commander Ready")
	fmt.Println("========================================")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.orchestrate(ctx)
		}
	}
}

func (c *Commander) orchestrate(ctx context.Context) {
	fmt.Printf("\n📊 Status: Progress %.1f%% | Status: %s\n", 
		c.state.OverallProgress, c.state.OverallStatus)

	for _, g := range c.guardians {
		if err := g.Execute(ctx, c.state); err != nil {
			log.Printf("❌ Guardian %s error: %v", g.Name(), err)
		}
	}
}

// ==================== Guardians ====================

type ScribeGuardian struct{}

func (ScribeGuardian) Name() string { return "L2-Scribe" }
func (ScribeGuardian) Execute(ctx context.Context, state *ThitnueahubState) error {
	state.Messages = append(state.Messages, "Daily report generated")
	fmt.Println("📝 L2 Scribe: Report updated")
	return nil
}

type AnalystGuardian struct{}

func (AnalystGuardian) Name() string { return "L4-Analyst" }
func (AnalystGuardian) Execute(ctx context.Context, state *ThitnueahubState) error {
	state.OverallProgress += 5.0
	fmt.Println("📊 L4 Analyst: Analysis completed")
	return nil
}

type BalancerGuardian struct{}

func (BalancerGuardian) Name() string { return "L11-Balancer" }
func (BalancerGuardian) Execute(ctx context.Context, state *ThitnueahubState) error {
	state.OverallStatus = "green"
	fmt.Println("⚖️ L11 Balancer: System balanced")
	return nil
}

// ==================== Main ====================

func main() {
	fmt.Println("🌟 TNH AI Orchestra v8.3 - 9th Era of Fire")
	fmt.Println("Private Fortress | Zero-Garbage Go Engine")

	commander := NewCommander()

	// Register Guardians
	commander.Register(ScribeGuardian{})
	commander.Register(AnalystGuardian{})
	commander.Register(BalancerGuardian{})

	// เพิ่ม Guardian อื่น ๆ ที่นี่

	commander.Run()
}
