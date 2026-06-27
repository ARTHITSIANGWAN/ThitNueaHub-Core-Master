package services

import (
	"errors"
	"net/http"
	"time"
)

// FetchData คือฟังก์ชันมาตรฐานที่มีความปลอดภัย
func FetchData(url string) (string, error) {
	// 1. หมวกกันน็อก (Safety Layer): ตรวจสอบเบื้องต้น
	if url == "" {
		return "", errors.New("invalid URL: Cannot process empty request")
	}

	// 2. การทำงาน (Operational Logic)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 3. ตรวจสอบการตอบกลับ (Auditing/FiveTwo Standard)
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("governance alert: API server returned non-200 status")
	}

	return "Data_Verified_Payload", nil
}
