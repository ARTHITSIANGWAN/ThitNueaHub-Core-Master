#!/data/data/com.termux/files/usr/bin/bash
# Project Name: TNH-AI-MASTER-TEMPLATE
# Folder Path: /
# File Name: update_daily_content.sh
cat << 'EOF' > cmd/video_builder/daily_rotator.go
package main

import (
"fmt"
"time"
)

type DailyTopic struct {
DayTheme string
Prompt   string
}

func GetTodayTopic() DailyTopic {
dayOfYear := time.Now().YearDay()
topics := []DailyTopic{
{"จริยธรรม AI และความโปร่งใส", "ใช้ AI อย่างมีศีลธรรม ไม่บิดเบือนข้อมูล ข้อควรระวังเรื่อง Deepfake"},
{"การป้อนข้อมูล (Prompt Engineering) ขั้นสูง", "วิธีสั่งการ AI ให้แม่นยำ ลดความผิดพลาด Zero-Garbage Prompting"},
{"การวิเคราะห์และตรวจสอบข้อเท็จจริง", "หลักคิดวิเคราะห์ แยกแยะ ก่อนแชร์ข้อมูลจาก AI สู่สาธารณะ"},
{"การสร้างคอนเทนต์ต้นทุนต่ำสำหรับ SME", "เทคนิคใช้ AI ช่วยธุรกิจฐานรากประหยัดเวลาและเพิ่มยอดขาย"},
}
return topics[dayOfYear % len(topics)]
}

func main() {
t := GetTodayTopic()
fmt.Printf("=== [TNH-AI DAILY ROTATOR] วันนี้ธีม: %s ===\\n", t.DayTheme)
fmt.Printf("โฟกัสเนื้อหา: %s\\n", t.Prompt)
}
EOF
go run cmd/video_builder/daily_rotator.go
