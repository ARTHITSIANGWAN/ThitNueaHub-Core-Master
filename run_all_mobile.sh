#!/data/data/com.termux/files/usr/bin/bash
# Project Name: TNH-AI-MASTER-TEMPLATE
# Folder Path: /
# File Name: run_all_mobile.sh
mkdir -p assets/video assets/audio output/scenes
echo ">>> [TNH-AI-MASTER-TEMPLATE] เริ่มรันระบบสร้างวิดีโอ 20 ซีน (3 นาที 20 วินาที)..."
if [ -f "cmd/video_builder/daily_rotator.go" ]; then
  go run cmd/video_builder/daily_rotator.go
else
  echo ">>> [INFO] กำลังเตรียมโครงสร้างระบบรายวัน..."
fi
