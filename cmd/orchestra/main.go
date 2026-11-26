// Project Name: TNH-AI-MASTER-TEMPLATE
// Folder Path: /cmd/orchestra/
// File Name: main.go

package main

import (
"encoding/json"
"fmt"
"log/slog"
"net/http"
"os"
"runtime"
"syscall"
"time"
)

type HealthResponse struct {
Node         string `json:"node"`
System       string `json:"system"`
Status       string `json:"status"`
GoAllocMB    string `json:"go_alloc_mb"`
FreeDiskGB   string `json:"free_disk_gb"`
Timestamp    string `json:"timestamp"`
}

func getFreeDiskGB(path string) string {
var stat syscall.Statfs_t
if err := syscall.Statfs(path, &stat); err != nil {
return "unknown"
}
freeBytes := stat.Bavail * uint64(stat.Bsize)
return fmt.Sprintf("%.2f GB", float64(freeBytes)/(1024*1024*1024))
}

func main() {
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

nodeName := os.Getenv("TNH_NODE_NAME")
if nodeName == "" {
nodeName = "OPPO-A79-5G-Primary"
}

mux := http.NewServeMux()

// 1. Endpoint รายงานสถานะเครื่องจริง (Real-Time Telemetry)
mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
var m runtime.MemStats
runtime.ReadMemStats(&m)

resp := HealthResponse{
Node:       nodeName,
System:     "TNH-AI-MASTER-TEMPLATE",
Status:     "operational",
GoAllocMB:  fmt.Sprintf("%.2f MB", float64(m.Alloc)/1024/1024),
FreeDiskGB: getFreeDiskGB("."),
Timestamp:  time.Now().Format(time.RFC3339),
}

w.Header().Set("Content-Type", "application/json")
w.Header().Set("Access-Control-Allow-Origin", "*")
_ = json.NewEncoder(w).Encode(resp)
})

// 2. เสิร์ฟหน้าเว็บ Transparent Glass (web/index.html) บนพอร์ต :2026
fs := http.FileServer(http.Dir("./web"))
mux.Handle("/", fs)

logger.Info("Pure Go Core & Web Dashboard Online", "port", 2026, "node", nodeName)
if err := http.ListenAndServe(":2026", mux); err != nil {
logger.Error("Server failed", "error", err)
os.Exit(1)
}
}
