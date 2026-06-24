#!/bin/bash

# ==============================================================================
# Project: ThitNueaHub | Premium 2000% Automation Deployment
# Core Context: Low Cost, High Intelligence (Zero-Garbage Mode)
# Authenticated System: narm-ing-agent (project-6e34f0b2)
# Dedicated To: Boss Art (Arthit)
# ==============================================================================

# --- [Configuration Variables] ---
# ข้อมูล Account ID และ Token ถอดจากสิทธิ์ระบบที่เจ้านายแนบมา
ACCOUNT_ID="58533a0fa29f04212ceb5c0a018e3412"
API_TOKEN="ใส่_API_TOKEN_Tnh_ai_ของเจ้านายตรงนี้"
SCRIPT_NAME="thitnueahub-core-worker"  # เปลี่ยนชื่อตาม Worker ที่ต้องการ
SOURCE_FILE="index.js"                 # ไฟล์โค้ดหลักของเจ้านาย

# --- [Validation Layer] ---
if [ ! -f "$SOURCE_FILE" ]; then
    echo "⚠️ [Error]: ไม่พบไฟล์ $SOURCE_FILE ในไดเรกทอรีนี้ค่ะเจ้านาย!"
    exit 1
fi

echo "🔔 [Context Nudge by kaew-ta] -> Sending code to nam-ing black-n backend..."
echo "💀 [nam-ing black-n]: Context Secured. Deploying to Cloudflare Global Network..."

# --- [Execution Layer - Zero Garbage HTTP PUT] ---
RESPONSE=$(curl -s -w "\n%{http_code}" -X PUT \
     "https://api.cloudflare.com/client/v4/accounts/${ACCOUNT_ID}/workers/scripts/${SCRIPT_NAME}" \
     -H "Authorization: Bearer ${API_TOKEN}" \
     -H "Content-Type: application/javascript" \
     --data-binary "@${SOURCE_FILE}")

# แยกผลลัพธ์เพื่อตรวจสอบความสะอาดของข้อมูล
HTTP_STATUS=$(echo "$RESPONSE" | tail -n 1)
BODY=$(echo "$RESPONSE" | sed '$d')

# --- [Response Evaluation] ---
if [ "$HTTP_STATUS" -eq 200 ]; then
    echo "✨ [Success]: Deploy เรียบร้อยแล้วค่ะเจ้านาย! ระบบออนไลน์ 100% (HTTP 200)"
else
    echo "❌ [Failed]: เกิดข้อผิดพลาดในการ Deploy (HTTP Status: ${HTTP_STATUS})"
    echo "💬 Response: ${BODY}"
fi
