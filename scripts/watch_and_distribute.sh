#!/data/data/com.termux/files/usr/bin/bash
# PROJECT NAME: TNH-AI-MASTER-TEMPLATE
# FOLDER PATH: /scripts/
# FILE NAME: watch_and_distribute.sh

# Project Name: TNH-AI-MASTER-TEMPLATE
# Folder Path: /scripts/
# File Name: watch_and_distribute.sh
WATCH_DIR="output"
TIKTOK_QUEUE="queue/tiktok"
YOUTUBE_QUEUE="queue/youtube"
mkdir -p "$WATCH_DIR" "$TIKTOK_QUEUE" "$YOUTUBE_QUEUE"
echo ">>> [TNH-AI PIPELINE] กำลังตรวจสอบโฟลเดอร์: $WATCH_DIR..."
for file in "$WATCH_DIR"/*.mp4; do
  if [ -f "$file" ]; then
    filename=$(basename "$file")
    echo ">>> พบไฟล์ใหม่: $filename -> กำลังกระจายเข้าคิวแพลตฟอร์ม..."
    cp "$file" "$TIKTOK_QUEUE/"
    cp "$file" "$YOUTUBE_QUEUE/"
    echo ">>> [DONE] ส่งไฟล์ $filename ไปยังคิว TikTok และ YouTube เรียบร้อย!"
  fi
done
