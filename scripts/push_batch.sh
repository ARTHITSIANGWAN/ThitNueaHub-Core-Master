#!/bin/bash
# PROJECT NAME: TNH-AI-MASTER-TEMPLATE
# FOLDER PATH: /scripts/
# FILE NAME: push_batch.sh

# Blueprint: ThitNueaHub Offline Batch Push Protocol (Port 22)

TARGET_DIR="/sdcard/ThitNueaHub_Local/tasks_by_group/eleven_white_tiger"
cd $TARGET_DIR

# 1. ตรวจสอบว่ามีข้อมูลอัปเดตออฟไลน์จากแอปในเครื่องไหม
if [ -n "$(git status --porcelain)" ]; then
    echo "[Jarvis] Found offline updates. Processing Zero-Garbage deployment..."
    
    # 2. จัดการแพ็กข้อมูลเข้าคิว
    git add .
    git commit -m "Offline Scheduled Sync: $(date '+%Y-%m-%d %H:%M:%S') [No Garbage]"
    
    # 3. เปิดท่อ Port 22 ยิงกระหน่ำก้อนเดียวจบ ทะลวงผ่าน SSH ด้วยสิทธิ์ narm-ing-agent
    echo "[Jarvis] Launching Port 22 Secure Pipeline..."
    ssh -p 22 git@github.com
    git push origin main
    
    echo "[Jarvis] Deployment successful. Cleaning loop memory."
else
    echo "[Jarvis] No offline data to speak. Standing by."
fi
