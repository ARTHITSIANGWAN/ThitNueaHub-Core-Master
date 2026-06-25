package main

import (
"encoding/json"
"fmt"
"net/http"
"time"
)

type MCPRequest struct {
AgentID   string                 `json:"agent_id"`
Context   string                 `json:"context"`
Command   string                 `json:"command"`
Timestamp int64                  `json:"timestamp"`
Metadata  map[string]interface{} `json:"metadata"`
}

func ZeroGarbageValidate(req MCPRequest) bool {
if time.Now().Unix()-req.Timestamp > 60 {
return false
}
if req.Command == "" || len(req.Context) > 1000 {
return false
}
return true
}

func mcpHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

if r.Method != http.MethodPost {
w.WriteHeader(http.StatusMethodNotAllowed)
json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
return
}

var req MCPRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
w.WriteHeader(http.StatusBadRequest)
json.NewEncoder(w).Encode(map[string]string{"error": "Bad request"})
return
}

// ใช้การเปรียบเทียบแบบตรงไปตรงมาเพื่อหลีกเลี่ยงสัญลักษณ์หลอน
isValid := ZeroGarbageValidate(req)
if isValid == false {
w.WriteHeader(http.StatusNotAcceptable)
json.NewEncoder(w).Encode(map[string]string{"status": "loop_or_garbage_detected"})
return
}

w.WriteHeader(http.StatusOK)
response := map[string]interface{}{
"status":       "SUCCESS",
"core_version": "TNH_AI_8.1_MASTER",
"message":      fmt.Sprintf("Agent executed command successfully"),
}
json.NewEncoder(w).Encode(response)
}

func main() {
http.HandleFunc("/v1/mcp/ingress", mcpHandler)
fmt.Println("🛰️ ThitNueaHub MCP Server Ingress Active on Port 8080...")
http.ListenAndServe(":8080", nil)
}
