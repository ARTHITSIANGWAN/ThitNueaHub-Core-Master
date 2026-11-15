#!/data/data/com.termux/files/usr/bin/bash
# PROJECT NAME: TNH-AI-MASTER-TEMPLATE
# FOLDER PATH: /scripts/
# FILE NAME: run_10s_video.sh

# Project Name: TNH-AI-MASTER-TEMPLATE
# Folder Path: /scripts/
# File Name: run_10s_video.sh
SUB="การใช้ภาษาที่ชัดเจนคือหัวใจสำคัญของการสื่อสารยุคเอไออย่างแท้จริงครับ"
TTS="การ ใช้ ภาษา ที่ ชัดเจน คือ หัวใจ สำคัญ ของ การ สื่วสาร ยุค เอไอ อย่าง แท้จริง ครับ"
echo ">>> [TNH-AI-MASTER-TEMPLATE] 15 Words / 10s Ready:"
echo "TTS: $TTS"
echo "SUB: $SUB"
if [ -f "raw_ai_output.mp4" ]; then
  ffmpeg -y -i raw_ai_output.mp4 -filter_complex "[0:v]setpts=PTS-STARTPTS,trim=duration=10[v];[0:a]atempo=0.98,apad=pad_dur=0.4,atrim=duration=10[a]" -map "[v]" -map "[a]" -c:v libx264 -preset fast -c:a aac -b:a 192k final_10s_100clear.mp4
  echo ">>> [DONE] Saved final_10s_100clear.mp4"
else
  echo ">>> [WAIT] รอไฟล์ raw_ai_output.mp4 สำหรับเรนเดอร์"
fi
