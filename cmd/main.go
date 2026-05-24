package main

import (
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
)

// Job represents the data structure for system tasks
type Job struct {
	JobID  string `json:"job_id"`
	Action string `json:"action"`
	Topic  string `json:"topic"`
	Ts     int64  `json:"ts"`
	Ttl    int    `json:"ttl"`
	Sig    string `json:"sig"`
}

// System State
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

	// Route definitions
	http.HandleFunc("/process", processHandler)
	http.HandleFunc("/metrics", metricsHandler)

	fmt.Printf("Engine initialization complete. Listening on %s\n", Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		panic(err)
	}
}

// worker processes jobs from the queue
func worker(id int) {
	for job := range jobQueue {
		fmt.Printf("Worker %d processing job: %s\n", id, job.JobID)
		
		// Update metrics
		metrics.Lock()
		metrics.Processed++
		metrics.Unlock()
	}
}

// processHandler accepts incoming job requests
func processHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var job Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Push job to the queue
	select {
	case jobQueue <- job:
		w.WriteHeader(http.StatusAccepted)
	default:
		http.Error(w, "Queue full", http.StatusServiceUnavailable)
	}
}

// metricsHandler reports engine performance
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics.Lock()
	defer metrics.Unlock()

	data, _ := json.Marshal(map[string]int{
		"processed_jobs": metrics.Processed,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
