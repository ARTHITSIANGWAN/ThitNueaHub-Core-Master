#!/data/data/com.termux/files/usr/bin/bash
# PROJECT NAME: TNH-AI-MASTER-TEMPLATE
# FOLDER PATH: /scripts/
# FILE NAME: auto_stream_24h.sh

# Project Name: TNH-AI-MASTER-TEMPLATE
# Folder Path: /scripts/
# File Name: auto_stream_24h.sh

set -e

BASE_DIR="$HOME/projects/TNH-AI-MASTER-TEMPLATE"
MEDIA_DIR="$BASE_DIR/media_loop"
mkdir -p "$MEDIA_DIR" "$BASE_DIR/queue/youtube" "$BASE_DIR/queue/tiktok"

echo "=== [TNH-AI] 24/7 Mobile Stream & Concat Engine ==="
echo "1) วางไฟล์ภาพปก (cover_16x9.png) และไฟล์เสียง (master_loop.mp3) ใน: $MEDIA_DIR"
echo "2) ตั้งค่าตัวแปร STREAM_KEY ก่อนเริ่มสตรีม (ห้ามเก็บ Key ในโค้ด)"

if [ -z "$STREAM_KEY" ]; then
  echo "⚠️ ยังไม่ได้ตั้งค่า STREAM_KEY (ตัวอย่าง: export STREAM_KEY='xxxx-xxxx-xxxx-xxxx')"
  exit 1
fi

# วนลูปสตรีมสด 24 ชม. อัตโนมัติ (หากเน็ตหลุดจะต่อใหม่ใน 5 วินาที)
while true; do
  echo "[$(date -Iseconds)] 📡 เริ่มส่งสัญญาณสตรีมสดไปยัง YouTube Live..."
  ffmpeg -re -stream_loop -1 -loop 1 -framerate 2 -i "$MEDIA_DIR/cover_16x9.png" \
    -stream_loop -1 -i "$MEDIA_DIR/master_loop.mp3" \
    -c:v libx264 -preset ultrafast -tune stillimage -pix_fmt yuv420p -g 4 -b:v 1500k \
    -c:a aac -b:a 128k -ar 44100 \
    -f flv "rtmp://a.rtmp.youtube.com/live2/$STREAM_KEY" || true
  echo "🔄 สัญญาณขาดช่วง กำลังเชื่อมต่อใหม่ใน 5 วินาที..."
  sleep 5
done
