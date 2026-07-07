#!/bin/bash
# 9thera Empire Security Guard - Lockdown Script
# ดำเนินการล็อกไฟล์สำคัญและกวาดล้างขยะตรรกะ

# พิกัดไฟล์ที่ต้องล็อก
LOCK_FILES=("baseline.json" "secret.key" "config.json" "templates/*.ml")

echo "[AUDIT] Starting Security Lockdown..."

# 1. เปลี่ยนสิทธิ์การเข้าถึงเป็น Read-Only (ป้องกันการแก้ไข)
for file in ${LOCK_FILES[@]}; do
    if [ -f $file ]; then
        chmod 400 $file
        chattr +i $file # สั่งล็อกถาวรในระดับ File System (ป้องกันแม้แต่ Root เปลี่ยนแปลง)
        echo "[LOCKED] $file is now Immutable."
    fi
done

# 2. กวาดล้างไฟล์ขยะและ Log เก่าที่อาจมี Secret หลุด
find . -name "*.log" -delete
find . -name "*.tmp" -delete
find . -name "*.bak" -delete

echo "[AUDIT] Lockdown Complete. Fortress Secured."

