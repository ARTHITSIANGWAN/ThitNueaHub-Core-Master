#!/bin/bash
# TNH_AI_DISTRIBUTOR: Synchronize Hub to Spoke
# Usage: ./scripts/distribute.sh [TARGET_FOLDER_NAME]

TARGET=$1

if [ -z "$TARGET" ]; then
    echo "❌ Error: ต้องระบุชื่อโปรเจกต์เป้าหมาย (เช่น TNH_AI_PROJECT_01)"
    exit 1
fi

SOURCE_DIR="."
TARGET_DIR="../$TARGET"

echo "🚀 กำลังดำเนินการ Sync: $SOURCE_DIR -> $TARGET_DIR"

# ใช้ rsync เพื่ออัปเดตไฟล์ (ข้าม .git เพราะเราต้องการแค่โครงสร้าง)
# --dry-run คือโหมด "ทดลอง" (ลบ --dry-run ออกเมื่อบอสพร้อมรันจริง)
rsync -av --progress \
    --exclude='.git' \
    --exclude='node_modules' \
    --exclude='dist' \
    --exclude='.env' \
    $SOURCE_DIR/ $TARGET_DIR/

echo "✅ Distribution Complete!"
