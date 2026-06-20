package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Configuration Constants
const (
	Port        = ":2026"
	QueueSize   = 128
	WorkerCount = 4
	SecretKey   = "THITNUEA_SECURE_KEY" // สำหรับ Sign ข้อมูลส่งขึ้น Edge
)

type Job struct {
	JobID  string `json:"job_id"`
	Action string `json:"action"`
	Topic  string `json:"topic"`
	Ts     int64  `json:"ts"`
	Ttl    int    `json:"ttl"`
	Sig    string `json:"sig"` // Signature เพื่อตรวจสอบความถูกต้อง
}

var (
	jobQueue = make(chan Job, QueueSize)
	metrics  = struct {
		sync.Mutex
		Processed int
	}{}
)

func main() {
	// Initialize worker pool
	for i := 0; i < WorkerCount; i++ {
		go worker(i)
	}

	http.HandleFunc("/process", processHandler)
	http.HandleFunc("/metrics", metricsHandler)

	fmt.Printf("Fortress Engine Ready on %s\n", Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		panic(err)
	}
}

// 🛡️ GATEKEEPER: ตรวจสอบสิทธิ์และ Topic
func isAllowed(topic string) bool {
	// Whitelist ตรวจสอบว่า Topic นี้อนุญาตให้ทำงานหรือไม่
	allowedTopics := map[string]bool{"CONTENT": true, "SYSTEM": true}
	return allowedTopics[topic]
}

// 🧠 PROCESSOR: เปลี่ยน Template ตาม DNA
func applyTemplate(action, topic string) ([]byte, error) {
	// Logic การเลือก Template (จำลองการอ่านไฟล์ .ml)
	content := fmt.Sprintf("Template Applied: %s for %s", action, topic)
	return []byte(content), nil
}

// 🌉 BRIDGE: ส่งข้อมูลขึ้น Edge อย่างปลอดภัย
func pushToEdge(data []byte, originalSig string) error {
	// ตรวจสอบ Signature ก่อนส่ง (กันข้อมูลปลอม)
	h := hmac.New(sha256.New, []byte(SecretKey))
	h.Write(data)
	calculatedSig := hex.EncodeToString(h.Sum(nil))

	if calculatedSig != originalSig {
		return fmt.Errorf("SECURITY ALERT: Signature mismatch")
	}

	// ส่งขึ้น Edge (Cloudflare)
	fmt.Printf("Bridge: Pushing data to Edge with Sig: %s\n", calculatedSig)
	return nil
}

// worker processes jobs
func worker(id int) {
	for job := range jobQueue {
		// 1. Gatekeeper Check
		if !isAllowed(job.Topic) {
			fmt.Printf("Worker %d: Rejected Topic %s\n", id, job.Topic)
			continue
		}

		// 2. Process Template
		payload, err := applyTemplate(job.Action, job.Topic)
		if err != nil {
			continue
		}

		// 3. Push to Bridge
		err = pushToEdge(payload, job.Sig)
		if err != nil {
			fmt.Printf("Worker %d: Bridge Error: %v\n", id, err)
			continue
		}

		metrics.Lock()
		metrics.Processed++
		metrics.Unlock()
	}
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var job Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	select {
	case jobQueue <- job:
		w.WriteHeader(http.StatusAccepted)
	default:
		http.Error(w, "Queue Full", http.StatusServiceUnavailable)
	}
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics.Lock()
	defer metrics.Unlock()
	fmt.Fprintf(w, "Processed: %d", metrics.Processed)
}
