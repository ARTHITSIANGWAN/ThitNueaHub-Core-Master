// PROJECT NAME: ThitNueaHub-Core-Master // FOLDER PATH: /cmd/security_audit/ // FILE NAME: main.go
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/ARTHITSIANGWAN/ThitNueaHub-Core-Master/backend/core"
)

var sqliteMagicHeader = []byte("SQLite format 3\x00")

func main() {
	dbPath := "tnh_core.db"
	if info, err := os.Stat(dbPath); err == nil && !info.IsDir() {
		f, err := os.Open(dbPath)
		if err != nil {
			log.Fatalf("open sqlite file failed: %v", err)
		}
		header := make([]byte, 16)
		_, readErr := f.Read(header)
		f.Close()
		if readErr != nil || !bytes.Equal(header, sqliteMagicHeader) {
			log.Fatalf("sqlite header audit failed for %s", dbPath)
		}
		fmt.Printf("SQLite Header Audit (%s): VALID (%d bytes)\n", dbPath, info.Size())
	} else {
		fmt.Printf("SQLite Header Audit: %s not initialized yet (skipped)\n", dbPath)
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

	fmt.Printf("Security & L8 HMAC Perimeter Audit (bind=%s): PASSED\n", core.GetBindAddr())
}
