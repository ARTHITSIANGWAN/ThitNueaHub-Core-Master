// Project Name: TNH-AI-MASTER-TEMPLATE
// Folder Path: /backend/pipeline/
// File Name: video_pacing_generator.go

package main

import (
"bytes"
"crypto/hmac"
"crypto/sha256"
"encoding/hex"
"encoding/json"
"fmt"
"net/http"
"os"
"os/exec"
"strings"
)

// [1] ค่าคงที่คุมจังหวะ: คลิป 10 วินาที จำกัดสูงสุด 15 คำ เพื่อให้พูดจบพอดีไม่ตัดคำท้าย
const (
TargetDurationSec = 10.0
MaxWordCount      = 15
GatewayEndpoint   = "http://127.0.0.1:2026/api/v1/media/generate"
)

// [2] โครงสร้าง Payload มาตรฐาน TNH-AI-MASTER-TEMPLATE
type VideoGenPayload struct {
Project        string  `json:"project"`
Model          string  `json:"model"`
Prompt         string  `json:"prompt"`
VoiceScript    string  `json:"voice_script"`
SubtitleText   string  `json:"subtitle_text"`
DurationSec    float64 `json:"duration_sec"`
AspectRatio    string  `json:"aspect_ratio"`
SpeakingRate   float64 `json:"speaking_rate"`
EnforceExact10 bool    `json:"enforce_exact_10"`
}

// [3] กรองคำพ้องเสียง (Phonetic Filter) แก้ปัญหา AI อ่านเพี้ยน 100%
func sanitizeThaiPhonetics(raw string) string {
replacer := strings.NewReplacer(
"สื่อ", "สื่ว",
"ตลาด", "ตะ-หลาด",
"สามารถ", "สา-มาด",
"ประวัติศาสตร์", "ประ-หวัด-ติ-สาด",
)
return replacer.Replace(raw)
}

// [4] ตรวจสอบโควตาคำศัพท์ (ห้ามเกิน 15 คำ สำหรับวิดีโอ 10 วินาที)
func validateWordBudget(words []string) error {
if len(words) > MaxWordCount {
return fmt.Errorf("คำเกินโควตา: มี %d คำ (จำกัดสูงสุด %d คำ ต่อ %.0f วินาที)", len(words), MaxWordCount, TargetDurationSec)
}
return nil
}

// [5] สร้างลายเซ็น HMAC-SHA256 ตามมาตรฐาน Security Guard ของ TNH-AI-MASTER-TEMPLATE
func signPayloadHMAC(data []byte, secret string) string {
h := hmac.New(sha256.New, []byte(secret))
h.Write(data)
return hex.EncodeToString(h.Sum(nil))
}

func main() {
// [6] สคริปต์ตัวอย่าง 15 คำเป๊ะ (เรื่องการใช้ภาษาและการสื่อสารยุค AI)
words := []string{
"การใช้", "ภาษา", "ที่", "ชัดเจน", "คือ",
"หัวใจ", "สำคัญ", "ของ", "การ", "สื่อสาร",
"ยุค", "เอไอ", "อย่าง", "แท้จริง", "ครับ",
}

// [7] ตรวจสอบจำนวนคำก่อนรัน
if err := validateWordBudget(words); err != nil {
fmt.Println("[ERROR]", err)
os.Exit(1)
}

// [8] แยกข้อความขึ้นจอ (ตัวสะกดถูก) กับ สคริปต์เสียงพากย์ (แปลงเสียงให้ AI อ่านชัด)
subtitleText := strings.Join(words, "")
ttsText := sanitizeThaiPhonetics(strings.Join(words, " "))

// [9] ประกอบร่าง Payload
payload := VideoGenPayload{
Project:        "TNH-AI-MASTER-TEMPLATE",
Model:          "veo-3-fast-generate",
Prompt:         "Cinematic dark studio, glassmorphism UI, clean Thai typography, 10 seconds exact, tnh'ai watermark",
VoiceScript:    ttsText,
SubtitleText:   subtitleText,
DurationSec:    TargetDurationSec,
AspectRatio:    "16:9",
SpeakingRate:   0.95,
EnforceExact10: true,
}

jsonData, _ := json.MarshalIndent(payload, "", "  ")
secretKey := os.Getenv("TNH_HMAC_SECRET")
if secretKey == "" {
secretKey = "tnh_master_default_key"
}
signature := signPayloadHMAC(jsonData, secretKey)

fmt.Println("=== [TNH-AI-MASTER-TEMPLATE] Video Payload (15 Words / 10s) ===")
fmt.Println(string(jsonData))
fmt.Println("X-TNH-Signature:", signature)

// [10] ยิงคำสั่งเข้า Fiber API Gateway (:2026) หากเปิดใช้งานอยู่
if os.Getenv("TNH_DISPATCH_LIVE") == "1" {
req, _ := http.NewRequest("POST", GatewayEndpoint, bytes.NewBuffer(jsonData))
req.Header.Set("Content-Type", "application/json")
req.Header.Set("X-TNH-Signature", signature)

client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
fmt.Println("[GATEWAY OFFLINE]", err)
} else {
defer resp.Body.Close()
fmt.Println("[DISPATCHED] Gateway Status:", resp.Status)
}
}

// [11] เตรียมคำสั่ง FFmpeg ล็อกความยาว 10.00 วินาทีเป๊ะ + เติมช่วงเงียบท้าย 0.3 วิ กันคำขาด
cmd := exec.Command("ffmpeg",
"-y",
"-i", "raw_ai_output.mp4",
"-af", "apad=pad_dur=0.3,atrim=0:10",
"-vf", "tpad=stop_mode=clone:stop_duration=0.3,trim=0:10,setpts=PTS-STARTPTS",
"-c:v", "libx264", "-preset", "fast", "-crf", "22",
"-c:a", "aac", "-b:a", "192k",
"final_10s_clear.mp4",
)
fmt.Println("\n=== Ready FFmpeg Lock Command ===")
fmt.Println(strings.Join(cmd.Args, " "))
}
