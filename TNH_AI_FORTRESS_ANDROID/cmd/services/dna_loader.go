package services

import (
	"fmt"
	"os"
)

// LoadIdentity อ่านข้อมูลจาก Context.ml เข้าสู่ระบบ
func LoadIdentity() string {
	content, err := os.ReadFile("templates/Context.ml")
	if err != nil {
		return "[WARNING] Identity DNA Not Found!"
	}
	// ในอนาคตเราจะเขียน Parser เพื่อแยกค่า ID ออกมา
	return string(content)
}
