// PROJECT NAME: ThitNueaHub-Core-Master 
// FOLDER PATH: /cmd/security_audit/ 
// FILE NAME: main.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ARTHITSIANGWAN/ThitNueaHub-Core-Master/backend/core"
)

func main() {
	dbPath := "tnh_core.db"
	if _, err := os.Stat(dbPath); err == nil {
		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			log.Fatalf("open sqlite failed: %v", err)
		}
		defer db.Close()

		var integrity string
		if err := db.QueryRow("PRAGMA integrity_check;").Scan(&integrity); err != nil {
			log.Fatalf("sqlite integrity check failed: %v", err)
		}
		fmt.Printf("SQLite Integrity (%s): %s\n", dbPath, integrity)
	} else {
		fmt.Printf("SQLite Audit: %s not initialized yet (skipped)\n", dbPath)
	}

	auditPayload := []byte("L1-CEO-001|L10-COMPLIANCE|AUDIT_PING")
	sig, err := core.SignPayload(auditPayload)
	if err != nil {
		log.Fatalf("L8 Glock audit sign failed: %v", err)
	}

	ok, err := core.VerifyPayload(auditPayload, sig)
	if err != nil || !ok {
		log.Fatalf("L8 Glock audit verify failed: %v", err)
	}

	fmt.Println("Security & L8 HMAC Perimeter Audit: PASSED")
}

