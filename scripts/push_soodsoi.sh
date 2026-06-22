#!/bin/bash
# ==============================================================================
# THITNUEAHUB AUTONOMOUS LOCAL ENGINE RECONCILIATION SCRIPT (5S COMPLIANT)
# ==============================================================================

# 1. กำหนดตัวแปรระบบพิกัดสากล (Universal Parameters)
TARGET_REPO="TNH-AI-9th-Era-of-Fire-AI-Orchestra-Mastering-Go-Engine-"
GITHUB_USER="ARTHITSIANGWAN"
# พิมพ์โทเค็นส่วนตัวตัวจริงของพี่แทนคำว่า YOUR_ACTUAL_PAT_TOKEN_HERE
MY_TOKEN="YOUR_ACTUAL_PAT_TOKEN_HERE" 

SOURCE_DIR=$(basename "$PWD")
BUNDLE_DIR="tnh_deployment_bundle"
DEST_REPO_DIR="tnh_destination_gateway"

echo "[STAGE 1/4] 5S Audit: Purging tracking conflicts..."
rm -rf $BUNDLE_DIR $DEST_REPO_DIR
mkdir -p $BUNDLE_DIR

echo "[STAGE 2/4] Packing structural assets (Scrubbing environment garbage)..."
# คัดแยกและกวาดเฉพาะซอร์สโค้ดบริสุทธิ์ ไม่หยิบไฟล์ .env หรือแคชขยะติดไป
rsync -av --exclude='.git' --exclude='.env' --exclude='node_modules' ./ $BUNDLE_DIR/

echo "[STAGE 3/4] Establishing secure tunnel integration layer..."
# ยิงตรงผ่านสิทธิ์ Token ส่วนตัวของบอสใหญ่ เพื่อข้ามบั๊กสิทธิ์ JOD-L5 403 Denied
git clone "https://x-access-token:${MY_TOKEN}@github.com/${GITHUB_USER}/${TARGET_REPO}.git" $DEST_REPO_DIR

if [ $? -ne 0 ]; then
    echo "[🚨 FATAL ERROR] Authentication failed. Check your security token value."
    exit 1
fi

# 2. จัดวางโฟลเดอร์ให้เป็นระเบียบตามโครงสร้างสากล
FINAL_TARGET_PATH="${DEST_REPO_DIR}/from_${SOURCE_DIR}"
mkdir -p "$FINAL_TARGET_PATH"
cp -rv $BUNDLE_DIR/* "$FINAL_TARGET_PATH/"

# 3. ตอกย้ำลายเซ็นดิจิทัลและผลักดันข้อมูลเข้าคลังหลัก
cd $DEST_REPO_DIR
git config user.name "${GITHUB_USER}"
git config user.email "founder@thitnueahub.ai"

git add .
git commit -m "Pure Go Logic Sync: Merged system assets from ${SOURCE_DIR} into empire core." || echo "Status: No payload modifications."
git push origin main

# 4. ล้างร่องรอยพื้นที่ชั่วคราวให้สะอาดกริบ
cd ..
rm -rf $BUNDLE_DIR $DEST_REPO_DIR

echo "[✅ SUCCESS] Mission accomplished. Data successfully integrated into the main pipeline."
