package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// ===== CONFIGURATION =====
var SECRET = []byte(getenv("SECRET", "9thera-empires-sovereign-secret-key-2026"))
const PORT = ":2026"
const QUEUE_SIZE = 128
const WORKERS = 4

// ===== DATA STRUCTURES =====
type Guardian struct {
	Level       string `json:"level"`
	Codename    string `json:"codename"`
	RoleTitle   string `json:"role_title"`
	WriteScope  string `json:"write_scope"`
	Description string `json:"description"`
}

type Job struct {
	JobID    string `json:"job_id"`
	Action   string `json:"action"`
	Topic    string `json:"topic"`
	Ts       int64  `json:"ts"`
	Ttl      int    `json:"ttl"`
	Sig      string `json:"sig,omitempty"`
}

type TraceConfig struct {
	Enabled          bool     `json:"enabled"`
	Destinations     []string `json:"destinations"`
	HeadSamplingRate  float64  `json:"head_sampling_rate"`
	Persist          bool     `json:"persist"`
}

type ObservabilityMatrix struct {
	Traces TraceConfig `json:"traces"`
	Logs   TraceConfig `json:"logs"`
}

// ===== GLOBAL STATE MATRIX =====
var (
	jobQueue       = make(chan Job, QUEUE_SIZE)
	scheduleKV     = sync.Map{} 
	codexMemories  = sync.Map{} 
	obsConfig      ObservabilityMatrix
	guardians      = []Guardian{
		{"L1", "ทิตเหนือ (Boss Art)", "Commander (CSO)", "root", "ผู้บัญชาการสูงสุด ถือสิทธิ์ราชา ควบคุมภาพรวมจักรวรรดิ"},
		{"L2", "แก้วตา", "Scribe (CDO)", "technical | vision_core", "หัวหน้าฝ่ายสื่อสาร สั่งรัน Training โมเดลภาพผ่าน Vertex AI คอยเก็บบอทภายนอก"},
		{"L3", "น้ำอิง", "Artist (CD)", "content | logic_recovery", "ขุนพลกู้คืนระบบและศัลยกรรมตรรกะระดับวิเคราะห์แยกแยะ n แบบมนุษย์"},
		{"L4", "พลายทอง", "Analyst (DIS)", "edge_intelligence", "นักวิเคราะห์เทคโนโลยีขอบสนาม ประมวลผลลอจิกบนอุปกรณ์เคลื่อนที่"},
		{"L5", "ไอ้จ๊อด", "Auditor (QA)", "technical_purge", "ผู้คุมกฎ Zero-Garbage สั่งยิงสคริปต์ [สุดซอย V3] ล้างขยะความจำทั้งหมด"},
		{"L6", "Merchant", "Merchant (BD)", "business_development", "แปลงศัพท์เทคนิคให้กลายเป็นโซลูชันเข้าถึงง่ายสำหรับกลุ่ม SME"},
		{"L7", "Oracle", "Oracle (SA)", "infrastructure_balance", "สถาปนิกระบบความเร็วแสง ผู้ออกแบบการไขว้พอร์ตและบาลานซ์ตู้เซฟ"},
		{"L8", "Glock", "Guardian (CISO)", "security_audit", "กล่องดำผู้พิทักษ์ สแกนสิทธิ์ตรวจสอบเอกสารความปลอดภัยสูงสุด"},
		{"L9", "Sentinel", "Sentinel", "leakage_prevention", "ผู้ตรวจการณ์เฝ้าระวังการรั่วไหลของข้อมูลและโมเดล AI ออกนอกอาณาจักร"},
		{"L10", "Auditor", "Auditor", "legal_compliance", "ผู้ตรวจสอบความเสี่ยงทางกฎหมาย ดักเช็คสิทธิ์การปฏิบัติตามกฎระเบียบ"},
		{"L11", "Balancer", "Balancer", "self_healing", "หน่วยบาลานซิ่งและชุบชีวิต Node หลังบ้าน เมื่อระบบติดหล่ม Error 13"},
	}
)

func init() {
	obsConfig = ObservabilityMatrix{
		Traces: TraceConfig{Enabled: true, Destinations: []string{"tnh-tracing-hub"}, HeadSamplingRate: 0.05, Persist: false},
		Logs:   TraceConfig{Enabled: true, Destinations: []string{"tnh-log-sink"}, HeadSamplingRate: 0.60, Persist: false},
	}
}

func main() {
	log.Println("⚡ [TNH V84.1 TRINITY CORE START]")
	for i := 1; i <= WORKERS; i++ { go startWorker(i) }

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<h1>🏰 TNH V84.1 MASTER CORE ACTIVE</h1><h3>Observability & Memories Locked</h3>")
	})

	mux.HandleFunc("/api/v84/trinity-status", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		duration := time.Since(startTime)
		actualLatency := fmt.Sprintf("%.4fms", float64(duration.Nanoseconds())/1e6)

		res := map[string]interface{}{
			"era": "9th_Era_of_Fire_Orchestra", "version": "V84.1_Observability_Hardened",
			"latency": actualLatency, "timestamp": time.Now(), "observability": obsConfig, "guardians": guardians,
		}
		_ = json.NewEncoder(w).Encode(res)
	})

	mux.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		var job Job
		if err := json.NewDecoder(r.Body).Decode(&job); err != nil { http.Error(w, "Bad Request", http.StatusBadRequest); return }
		job.Ts = time.Now().Unix(); job.Ttl = 120; job.Sig = generateSignature(job)
		img := image.NewRGBA(image.Rect(0, 0, 10, 10))
		jobBytes, _ := json.Marshal(job)
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("X-TNH-Payload", hex.EncodeToString(jobBytes))
		_ = png.Encode(w, img)
	})

	mux.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		fmt.Fprintf(w, "🔍 [Decoder]: พ่นสัจจะข้อมูลสำเร็จ -> %s\n", string(body))
	})

	mux.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		var job Job
		if err := json.NewDecoder(r.Body).Decode(&job); err != nil { http.Error(w, "Invalid Payload", http.StatusBadRequest); return }
		if !verifySignature(job, job.Sig) { http.Error(w, "🛡️ Signature ผิดพลาด!", http.StatusUnauthorized); return }
		select {
		case jobQueue <- job:
			w.WriteHeader(http.StatusAccepted)
			_, _ = w.Write([]byte(`{"status":"queued"}`))
		default: http.Error(w, "Queue Full", http.StatusServiceUnavailable)
		}
	})

	mux.HandleFunc("/schedule", func(w http.ResponseWriter, r *http.Request) {
		var job Job
		if err := json.NewDecoder(r.Body).Decode(&job); err != nil { http.Error(w, "Bad Request", http.StatusBadRequest); return }
		scheduleKV.Store(job.JobID, job)
		codexMemories.Store(fmt.Sprintf("mem_%s", job.JobID), job.Topic)
		log.Printf("📅 [Memory Matrix]: ล็อกความจำงาน %s สำเร็จ\n", job.JobID)
		_, _ = w.Write([]byte(`{"status":"scheduled_and_memory_locked"}`))
	})

	log.Printf("👑 MASTER TEMPLATE ENGINE V84.1 ONLINE | Port: %s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}

func generateSignature(j Job) string {
	h := hmac.New(sha256.New, SECRET)
	h.Write([]byte(fmt.Sprintf("%s:%s:%s:%d", j.JobID, j.Action, j.Topic, j.Ts)))
	return hex.EncodeToString(h.Sum(nil))
}

func verifySignature(j Job, sig string) bool { return generateSignature(j) == sig }
func startWorker(id int) { for job := range jobQueue { log.Printf("🤖 [ขุนพลเสือขาว %d]: ลุยงาน ID: %s (%s)\n", id, job.JobID, job.Topic) } }
func getenv(key, fallback string) string { if v, exists := os.LookupEnv(key); exists { return v }; return fallback }
