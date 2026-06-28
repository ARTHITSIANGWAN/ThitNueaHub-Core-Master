package core

import (
	"context"
	"fmt"
	"log"
	"time"

	"thitnueahub/internal/guardians"
)

type ThitnueahubState struct {
	Messages        []string
	OverallProgress float64
	OverallStatus   string
	ActiveIssues    []string
	DebateHistory   []string
	GuardianStatus  map[string]string
}

type Orchestrator struct {
	guardians []guardians.Guardian
	state     *ThitnueahubState
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		guardians: make([]guardians.Guardian, 0),
		state: &ThitnueahubState{
			GuardianStatus: make(map[string]string),
		},
	}
}

func (o *Orchestrator) RegisterAll() {
	o.register(guardians.NewScribe())
	o.register(guardians.NewAnalyst())
	o.register(guardians.NewBalancer())
	fmt.Println("✅ All Guardians Registered")
}

func (o *Orchestrator) register(g guardians.Guardian) {
	o.guardians = append(o.guardians, g)
	o.state.GuardianStatus[g.Name()] = "registered"
}

func (o *Orchestrator) Start() {
	fmt.Println("🚀 Orchestra Running...")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		o.orchestrate(context.Background())
	}
}

func (o *Orchestrator) orchestrate(ctx context.Context) {
	fmt.Printf("\n📊 Status: %.1f%% | %s\n", o.state.OverallProgress, o.state.OverallStatus)

	for _, g := range o.guardians {
		if err := g.Execute(ctx, o.state); err != nil {
			log.Printf("❌ %s error: %v", g.Name(), err)
		}
	}
}
