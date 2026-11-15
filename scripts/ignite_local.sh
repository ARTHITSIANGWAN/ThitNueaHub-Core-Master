#!/bin/bash
# PROJECT NAME: TNH-AI-MASTER-TEMPLATE
# FOLDER PATH: /scripts/
# FILE NAME: ignite_local.sh

# TNH LOCAL WORKFLOW EMULATOR - PURE BASH LOGIC

TARGET_REPO="TNH-AI-9th-Era-of-Fire-AI-Orchestra-Mastering-Go-Engine-"
DEST_DIR="local_sync_destination"

echo "[1/3] 5S Audit: Cleaning transient data artifacts..."
rm -rf $DEST_DIR temp_bundle
mkdir -p temp_bundle

echo "[2/3] Packing target assets cleanly (Filtering local garbage)..."
# กวาดข้อมูลเฉพาะไฟล์งานสากล ไม่หยิบ .env และคุกกี้ที่ค้างในเครื่อง
rsync -av --exclude='.git' --exclude='.env' --exclude='node_modules' ./ temp_bundle/

echo "[3/3] Commencing connection setup to remote empire repository..."
# ทำการดึงคลังปลายทางลงมาจัดการในพื้นที่แยกส่วน
git clone "https://github.com/ARTHITSIANGWAN/$TARGET_REPO.git" $DEST_DIR

# ทำการย้ายของลงเลนเฉพาะโปรเจกลูก
FINAL_DEST="$DEST_DIR/from_TNH_AI_Master_Template"
mkdir -p "$FINAL_DEST"
cp -rv temp_bundle/* "$FINAL_DEST/"

cd $DEST_DIR
git add .
git commit -m "Level 5 Local Terminal Sync: Core payload integrated cleanly."
git push origin main

echo "[✅ SUCCESS] Local deployment sequence finished successfully."
