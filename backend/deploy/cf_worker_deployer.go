// PROJECT NAME: TNH-AI-MASTER-TEMPLATE [TNH-MT]
// FOLDER PATH: /backend/deploy/
// FILE NAME: cf_worker_deployer.go
// ADDRESS: SC/MB
// DAY: 18/11

package main

import (
"bufio"
"bytes"
"fmt"
"io"
"mime/multipart"
"net/http"
"net/textproto"
"os"
"path/filepath"
"strings"
"time"
)

func loadEnv(envPath string) {
f, err := os.Open(envPath)
if err != nil {
return
}
defer f.Close()
scanner := bufio.NewScanner(f)
for scanner.Scan() {
line := strings.TrimSpace(scanner.Text())
if line == "" || strings.HasPrefix(line, "#") {
continue
}
parts := strings.SplitN(line, "=", 2)
if len(parts) == 2 {
k := strings.TrimSpace(parts[0])
v := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
if os.Getenv(k) == "" {
_ = os.Setenv(k, v)
}
}
}
}

func main() {
homeDir, _ := os.UserHomeDir()
projDir := filepath.Join(homeDir, "TNH-AI-MASTER-TEMPLATE")
loadEnv(filepath.Join(projDir, ".env"))

accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
workerName := os.Getenv("CLOUDFLARE_WORKER_NAME")
if workerName == "" {
workerName = "tnh"
}

if accountID == "" || apiToken == "" || strings.Contains(apiToken, "วาง_TOKEN") {
fmt.Println("❌ กรุณาใส่ CLOUDFLARE_API_TOKEN จริงใน ~/TNH-AI-MASTER-TEMPLATE/.env ก่อนสั่ง Deploy")
os.Exit(1)
}

scriptBytes, err := os.ReadFile(filepath.Join(projDir, "worker", "tnh_edge_worker.js"))
if err != nil {
fmt.Printf("❌ อ่านไฟล์ tnh_edge_worker.js ไม่ได้: %v\n", err)
os.Exit(1)
}

var body bytes.Buffer
writer := multipart.NewWriter(&body)

metaHeader := make(textproto.MIMEHeader)
metaHeader.Set("Content-Disposition", `form-data; name="metadata"`)
metaHeader.Set("Content-Type", "application/json")
metaPart, _ := writer.CreatePart(metaHeader)
_, _ = metaPart.Write([]byte(`{"main_module":"tnh_edge_worker.js","compatibility_date":"2026-11-01"}`))

jsHeader := make(textproto.MIMEHeader)
jsHeader.Set("Content-Disposition", `form-data; name="tnh_edge_worker.js"; filename="tnh_edge_worker.js"`)
jsHeader.Set("Content-Type", "application/javascript+module")
jsPart, _ := writer.CreatePart(jsHeader)
_, _ = jsPart.Write(scriptBytes)
_ = writer.Close()

apiURL := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/workers/scripts/%s", accountID, workerName)
req, _ := http.NewRequest(http.MethodPut, apiURL, &body)
req.Header.Set("Authorization", "Bearer "+apiToken)
req.Header.Set("Content-Type", writer.FormDataContentType())

client := &http.Client{Timeout: 30 * time.Second}
resp, err := client.Do(req)
if err != nil {
fmt.Printf("❌ เชื่อมต่อ Cloudflare API ล้มเหลว: %v\n", err)
os.Exit(1)
}
defer resp.Body.Close()
respBody, _ := io.ReadAll(resp.Body)

if resp.StatusCode >= 200 && resp.StatusCode < 300 {
fmt.Printf("✅ Deploy Worker '%s' ขึ้น Cloudflare Edge สำเร็จ (HTTP %d)\n", workerName, resp.StatusCode)
} else {
fmt.Printf("❌ Cloudflare API Error (HTTP %d): %s\n", resp.StatusCode, string(respBody))
os.Exit(1)
}
}
