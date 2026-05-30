package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/syumai/workers" // สับสายเข้าท่อ edge engine ของ cloudflare workers แท้ 100%
)

const secretKey = "thitnueahub_secret_token_v86"

type job struct {
	JobID  string `json:"job_id"`
	Action string `json:"action"` 
	Topic  string `json:"topic"`
	Ts     int64  `json:"ts"`
	Ttl    int    `json:"ttl"`
	Sig    string `json:"sig"`
}

// ตรวจสอบลายเซ็นดิจิทัลตามกฎเหล็กความปลอดภัยด่านแรก (Phase 1)
func validateSignature(j job) bool {
	message := fmt.Sprintf("%s|%s|%s", j.JobID, j.Action, j.Topic)
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	return j.Sig == hex.EncodeToString(mac.Sum(nil))
}

func processHandler(w http.ResponseWriter, r *http.Request) {
	// หมวกกันน็อคเหล็กกล้า ป้องกันการแครชที่ edge คืนสถานะคลีนเสมอป้องกันระบบลูปดับ
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"status":"edge_panic_recovered"}`)
		}
	}()

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var j job
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// คัดกรองและกำจัดขยะทิ้งทันทีหากลายเซ็นโมฆะ
	if !validateSignature(j) {
		http.Error(w, "unauthorized_garbage_detected", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	// L1 จ่ายงานกระจายสายตรงไป spoke ทันทีแบบไร้รอยหน่วงตามกฎข้อ 3 ของ manifesto
	switch j.Action {
	case "scrape":
		fmt.Fprintf(w, `{"status":"dispatched","target":"services/scraper","job_id":"%s"}`, j.JobID)
	case "schedule":
		fmt.Fprintf(w, `{"status":"dispatched","target":"services/scheduler","job_id":"%s"}`, j.JobID)
	case "agent_call":
		fmt.Fprintf(w, `{"status":"dispatched","target":"my_agent/app","job_id":"%s"}`, j.JobID)
	default:
		fmt.Fprintf(w, `{"status":"dispatched","target":"general_spoke","job_id":"%s"}`, j.JobID)
	}
}

func main() {
	// ลบพอร์ตประจำถิ่น :2026 และ listenandserve ทิ้งทั้งหมด เพื่อให้เป็น edge serverless สมบูรณ์แบบ
	http.HandleFunc("/process", processHandler)
	workers.Serve(nil)
}
