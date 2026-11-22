// PROJECT NAME: ThitNueaHub-Core-Master // FOLDER PATH: /backend/core/ // FILE NAME: config_addr.go
package core

import "os"

// GetBindAddr returns the gateway listen address from env with a safe localhost default.
func GetBindAddr() string {
	if addr := os.Getenv("TNH_BIND_ADDR"); addr != "" {
		return addr
	}
	return "127.0.0.1:2026"
}
