// PROJECT NAME: ThitNueaHub-Core-Master // FOLDER PATH: /cmd/envelope_demo/ // FILE NAME: main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ARTHITSIANGWAN/ThitNueaHub-Core-Master/backend/core"
)

type LayerIdentity struct {
	AgentID     string    `json:"agent_id"`
	Codename    string    `json:"codename"`
	Designation string    `json:"designation"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type CommandEnvelope struct {
	SourceLayer string `json:"source_layer"`
	TargetLayer string `json:"target_layer"`
	Action      string `json:"action"`
	Signature   string `json:"signature"`
}

func main() {
	ceo := LayerIdentity{
		AgentID:     "L1-CEO-001",
		Codename:    "ThitNuea",
		Designation: "CEO",
		Status:      "ACTIVE",
		CreatedAt:   time.Now().UTC(),
	}

	identityJSON, err := json.Marshal(ceo)
	if err != nil {
		log.Fatalf("marshal identity failed: %v", err)
	}
	fmt.Printf("L1 Identity Ready: %s\n", string(identityJSON))

	action := "EXECUTE_5S_CLEAN_AUDIT"
	sig, err := core.SignPayload([]byte(action))
	if err != nil {
		log.Fatalf("L8 Glock signing failed (check TNH_MASTER_SECRET): %v", err)
	}

	envelope := CommandEnvelope{
		SourceLayer: "L1-CEO",
		TargetLayer: "L5-ENGINE",
		Action:      action,
		Signature:   sig,
	}

	valid, err := core.VerifyPayload([]byte(envelope.Action), envelope.Signature)
	if err != nil || !valid {
		log.Fatalf("L8 Glock verification rejected envelope: %v", err)
	}

	fmt.Printf("Verified Envelope [%s -> %s]: bind=%s action=%s\n",
		envelope.SourceLayer, envelope.TargetLayer, core.GetBindAddr(), envelope.Action)
}
