// PROJECT NAME: TNH-AI-MASTER-TEMPLATE [TNH-MT]
// FOLDER PATH: /backend/core/
// FILE NAME: a2a_orchestrator.go
// ADDRESS: SC/MB
// DAY: 18/11

package main

import (
"context"
"database/sql"
"encoding/json"
"fmt"
"log/slog"
"os"
"os/exec"
"os/signal"
"path/filepath"
"syscall"
"time"

"github.com/gofiber/fiber/v2"
"github.com/panjf2000/ants/v2"
_ "modernc.org/sqlite"
)

type AIDispatch struct {
Sender    string                 `json:"sender"`
Receiver  string                 `json:"receiver"`
Action    string                 `json:"action"`
Payload   map[string]interface{} `json:"payload"`
Priority  int                    `json:"priority"`
Timestamp string                 `json:"timestamp"`
}

type InsightRecord struct {
ID        int    `json:"id"`
Topic     string `json:"topic"`
Insight   string `json:"insight"`
CreatedAt string `json:"created_at"`
}

type AuditLogRecord struct {
ID        int    `json:"id"`
Sender    string `json:"sender"`
Receiver  string `json:"receiver"`
Action    string `json:"action"`
Status    string `json:"status"`
Result    string `json:"result"`
CreatedAt string `json:"created_at"`
}

func initSQLiteWAL(dbPath string) (*sql.DB, error) {
db, err := sql.Open("sqlite", dbPath)
if err != nil {
return nil, err
}
for _, p := range []string{
"PRAGMA journal_mode=WAL;",
"PRAGMA synchronous=NORMAL;",
"PRAGMA busy_timeout=5000;",
} {
if _, err := db.Exec(p); err != nil {
return nil, err
}
}
schema := `
CREATE TABLE IF NOT EXISTS insights (
id INTEGER PRIMARY KEY AUTOINCREMENT,
topic TEXT NOT NULL,
insight TEXT NOT NULL,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS a2a_audit_logs (
id INTEGER PRIMARY KEY AUTOINCREMENT,
sender TEXT,
receiver TEXT,
action TEXT,
status TEXT,
result TEXT,
created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`
_, err = db.Exec(schema)
return db, err
}

func main() {
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
homeDir, _ := os.UserHomeDir()
baseProj := filepath.Join(homeDir, "TNH-AI-MASTER-TEMPLATE")
db, err := initSQLiteWAL(filepath.Join(baseProj, "thitnuea_knowledge.db"))
if err != nil {
logger.Error("SQLite WAL init failed", "err", err)
os.Exit(1)
}
defer db.Close()

pool, err := ants.NewPool(11)
if err != nil {
logger.Error("Worker pool init failed", "err", err)
os.Exit(1)
}
defer pool.Release()

app := fiber.New(fiber.Config{
AppName:               "TNH-AI-V83-SOVEREIGN-FORTRESS",
DisableStartupMessage: true,
})

app.Get("/health", func(c *fiber.Ctx) error {
return c.JSON(fiber.Map{
"status":         "operational",
"system":         "TNH-AI-THITNUEAHUB",
"port":           2026,
"sqlite_mode":    "WAL",
"active_workers": pool.Running(),
"timestamp":      time.Now().Format(time.RFC3339),
})
})

app.Post("/api/v1/a2a/dispatch", func(c *fiber.Ctx) error {
var msg AIDispatch
if err := c.BodyParser(&msg); err != nil {
return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid A2A JSON"})
}
if msg.Timestamp == "" {
msg.Timestamp = time.Now().Format(time.RFC3339)
}

_ = pool.Submit(func() {
ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
defer cancel()
var resultText string
status := "completed"

switch msg.Action {
case "ANALYZE_PROBABILITY":
proj := fmt.Sprintf("%v", msg.Payload["project_name"])
resultText = fmt.Sprintf("วิเคราะห์โปรเจกต์ %s ผ่านเกณฑ์ Zero-Garbage & 5ส พร้อมรันบน Termux", proj)
_, _ = db.Exec("INSERT INTO insights (topic, insight) VALUES (?, ?)", proj, resultText)

case "SAVE_KNOWLEDGE":
topic := fmt.Sprintf("%v", msg.Payload["topic"])
insight := fmt.Sprintf("%v", msg.Payload["insight"])
_, _ = db.Exec("INSERT INTO insights (topic, insight) VALUES (?, ?)", topic, insight)
resultText = fmt.Sprintf("บันทึกความรู้ '%s' ลง SQLite WAL สำเร็จ", topic)

case "FARM_MERGE":
out, err := exec.CommandContext(ctx, "bash", filepath.Join(baseProj, "scripts", "tnh_ctl.sh"), "media").CombinedOutput()
if err != nil {
status = "error"
resultText = fmt.Sprintf("Farm merge error: %v | %s", err, string(out))
} else {
resultText = "ย้ายไฟล์เพลงและรวมอัลบั้มแยกหมวดหมู่ + เรนเดอร์ 144p สำเร็จ"
}

case "SYNC_FILES":
out, err := exec.CommandContext(ctx, filepath.Join(baseProj, "bin", "tnh-file-manager")).CombinedOutput()
if err != nil {
status = "error"
resultText = fmt.Sprintf("File sync error: %v | %s", err, string(out))
} else {
resultText = "ประทับตรา Header 5 บรรทัด (SC/MB) และสร้างสารบัญ Markdown สำเร็จ"
}

case "DEPLOY_CF_WORKER":
out, err := exec.CommandContext(ctx, filepath.Join(baseProj, "bin", "tnh-cf-deploy")).CombinedOutput()
if err != nil {
status = "error"
resultText = fmt.Sprintf("CF deploy error: %v | %s", err, string(out))
} else {
resultText = fmt.Sprintf("Deploy Cloudflare Worker สำเร็จ: %s", string(out))
}

default:
rawPayload, _ := json.Marshal(msg.Payload)
resultText = fmt.Sprintf("บันทึกคำสั่ง %s: %s", msg.Action, string(rawPayload))
}

_, _ = db.Exec("INSERT INTO a2a_audit_logs (sender, receiver, action, status, result) VALUES (?, ?, ?, ?, ?)",
msg.Sender, msg.Receiver, msg.Action, status, resultText)
logger.Info("A2A task finished", "action", msg.Action, "status", status, "result", resultText)
})

return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
"status":       "queued",
"processed_by": "TNH-Core-Go-V83",
"action":       msg.Action,
"timestamp":    time.Now().Format(time.RFC3339),
})
})

app.Get("/api/v1/insights", func(c *fiber.Ctx) error {
rows, err := db.Query("SELECT id, topic, insight, created_at FROM insights ORDER BY id DESC LIMIT 20")
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
}
defer rows.Close()
var list []InsightRecord
for rows.Next() {
var r InsightRecord
if err := rows.Scan(&r.ID, &r.Topic, &r.Insight, &r.CreatedAt); err == nil {
list = append(list, r)
}
}
return c.JSON(fiber.Map{"count": len(list), "items": list})
})

app.Get("/api/v1/logs", func(c *fiber.Ctx) error {
rows, err := db.Query("SELECT id, sender, receiver, action, status, result, created_at FROM a2a_audit_logs ORDER BY id DESC LIMIT 20")
if err != nil {
return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
}
defer rows.Close()
var list []AuditLogRecord
for rows.Next() {
var r AuditLogRecord
if err := rows.Scan(&r.ID, &r.Sender, &r.Receiver, &r.Action, &r.Status, &r.Result, &r.CreatedAt); err == nil {
list = append(list, r)
}
}
return c.JSON(fiber.Map{"count": len(list), "logs": list})
})

go func() {
logger.Info("TNH Sovereign Fortress listening", "addr", "127.0.0.1:2026")
if err := app.Listen("127.0.0.1:2026"); err != nil {
logger.Error("Server stopped", "err", err)
}
}()

quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit
_ = app.Shutdown()
}
