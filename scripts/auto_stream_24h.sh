#!/data/data/com.termux/files/usr/bin/bash
# PROJECT NAME: TNH-AI-THITNUEAHUB
# FOLDER PATH: /scripts/
# FILE NAME: auto_stream_24h.sh

set -u

BASE_DIR="$HOME/projects/TNH-AI-MASTER-TEMPLATE"
VAULT_FILE="$HOME/.tnh_vault.env"
PLAYLIST_FILE="$BASE_DIR/media/playlist.txt"
LOG_FILE="$BASE_DIR/logs/stream_24h.log"
CYCLE_SECONDS=41400 # 11.5 ชั่วโมง (ป้องกันเกินลิมิต 12 ชม. ของ YouTube Archive)

# ปลดล็อก Wake Lock อัตโนมัติเมื่อสั่งหยุดสคริปต์
cleanup() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] หยุดการทำงานสตรีม และปลดล็อก Wake Lock" | tee -a "$LOG_FILE"
  command -v termux-wake-unlock >/dev/null 2>&1 && termux-wake-unlock
  exit 0
}
trap cleanup SIGINT SIGTERM

# 1. เปิด Wake Lock กัน Android ตัดโปรเซสตอนปิดหน้าจอ
if command -v termux-wake-lock >/dev/null 2>&1; then
  termux-wake-lock
fi

# 2. โหลด Stream Key จากตู้เซฟ (ห้ามฮาร์ดโค้ดในสคริปต์)
if [ -f "$VAULT_FILE" ]; then
  # shellcheck disable=SC1090
  source "$VAULT_FILE"
else
  echo "❌ ไม่พบไฟล์ $VAULT_FILE กรุณาสร้างไฟล์เก็บคีย์ก่อนรัน" | tee -a "$LOG_FILE"
  exit 1
fi

if [ -z "${YOUTUBE_STREAM_KEY:-}" ] || [ "$YOUTUBE_STREAM_KEY" = "ใส่_STREAM_KEY_ของคุณที่นี่" ]; then
  echo "❌ กรุณาใส่ YOUTUBE_STREAM_KEY จริงใน $VAULT_FILE ก่อนเริ่มสตรีม" | tee -a "$LOG_FILE"
  exit 1
fi

# 3. ตรวจสอบไฟล์รายการเพลง/วิดีโอ (playlist.txt)
if [ ! -f "$PLAYLIST_FILE" ]; then
  echo "❌ ไม่พบไฟล์ $PLAYLIST_FILE (ตัวอย่างรูปแบบในไฟล์: file '/data/data/com.termux/files/home/projects/TNH-AI-MASTER-TEMPLATE/media/track01.mp4')" | tee -a "$LOG_FILE"
  exit 1
fi

RTMP_URL="rtmp://a.rtmp.youtube.com/live2/${YOUTUBE_STREAM_KEY}"
ROUND=1

echo "🚀 [TNH-AI-THITNUEAHUB] เริ่มระบบสตรีมสด 24 ชม. (โหมดประหยัด CPU -c copy)" | tee -a "$LOG_FILE"

# 4. ลูปสตรีมต่อเนื่อง ตัดรอบทุก 11.5 ชั่วโมง
while true; do
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] ▶️ เริ่มสตรีมรอบที่ $ROUND (จำกัดเวลา 11.5 ชม. / ${CYCLE_SECONDS}s)" | tee -a "$LOG_FILE"

  # ใช้ timeout คุมชั้นนอกเผื่อเน็ตค้าง + ใช้ -t 41400 ปิดสตรีมอย่างนุ่มนวล
  timeout -k 15 $((CYCLE_SECONDS + 60)) ffmpeg \
    -hide_banner \
    -loglevel warning \
    -re \
    -stream_loop -1 \
    -f concat \
    -safe 0 \
    -i "$PLAYLIST_FILE" \
    -c copy \
    -t "$CYCLE_SECONDS" \
    -f flv "$RTMP_URL" >> "$LOG_FILE" 2>&1

  EXIT_CODE=$?
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] ⏹️ จบรอบที่ $ROUND (Exit Code: $EXIT_CODE) — พัก 8 วินาทีให้ YouTube บันทึกย้อนหลัง..." | tee -a "$LOG_FILE"
  
  ROUND=$((ROUND + 1))
  sleep 8
done
