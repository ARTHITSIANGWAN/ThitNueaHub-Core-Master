#!/bin/bash
# ==============================================================================
# THITNUEAHUB AUTONOMOUS LOCAL ENGINE RECONCILIATION SCRIPT (5S COMPLIANT)
# ==============================================================================

# 1. ปรับปรุงพิกัดสัจข้อมูลให้ตรงตามระบบคลังหลัก v83 จริง
TARGET_REPO="TNH-AI-V83-TRINITY-EMPIRE"
GITHUB_USER="ARTHITSIANGWAN"
MY_TOKEN="ghp_xxxxYOUR_REAL_GITHUB_TOKENxxxx" # 💡 แก้ไข: ใส่ token แท้ของแอดมินตรงนี้

BUNDLE_DIR="tnh_deployment_bundle"
DEST_REPO_DIR="tnh_destination_gateway"

echo "[STAGE 1/4] 5S Audit: Purging tracking conflicts..."
rm -rf $BUNDLE_DIR $DEST_REPO_DIR
mkdir -p $BUNDLE_DIR

echo "[STAGE 2/4] Packing structural assets..."
# คัดแยกเฉพาะซอร์สโค้ดบริสุทธิ์ ไม่เอาไฟล์ขยะติดไป
rsync -av --exclude='.git' --exclude='.env' --exclude='node_modules' ./ $BUNDLE_DIR/

echo "[STAGE 3/4] Establishing secure tunnel integration layer..."
git clone "https://x-access-token:${MY_TOKEN}@github.com/${GITHUB_USER}/${TARGET_REPO}.git" $DEST_REPO_DIR

if [ $? -ne 0 ]; then
    echo "[🚨 FATAL ERROR] Authentication failed. Check your security token value."
    exit 1
fi

# 2. 💡 ศัลยกรรม: วางไฟล์ลงที่ root ของคลังหลักโดยตรง เพื่อให้น่านฟ้า cloudflare ทำงานได้ 100%
cp -rv $BUNDLE_DIR/* $DEST_REPO_DIR/

# 3. ตอกย้ำลายเซ็นดิจิทัลและผลักดันข้อมูลเข้าคลังหลัก
cd $DEST_REPO_DIR
git config user.name "${GITHUB_USER}"
git config user.email "founder@thitnueahub.ai"

git add .
git commit -m "pure go & typescript logic sync: architecture unified 100%." || echo "status: no payload modifications."
git push origin main

# 4. ล้างร่องรอยพื้นที่ชั่วคราวให้สะอาดกริบไร้ขยะจราจร
cd ..
rm -rf $BUNDLE_DIR $DEST_REPO_DIR

echo "[✅ SUCCESS] Mission accomplished. Data integrated into root pipeline."
