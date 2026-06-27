#!/bin/bash
# TNH_AI_FORTRESS_ANDROID Setup Script

echo "TNH_AI_FORTRESS_ANDROID"

# สร้างโครงสร้างหลัก
mkdir -p cmd docs services/scraper services/scheduler data/reports data/content_stock templates scripts

# สร้างไฟล์ Placeholder เพื่อให้ Git มองเห็นโฟลเดอร์
touch docs/.gitkeep services/scraper/.gitkeep services/scheduler/.gitkeep data/reports/.gitkeep data/content_stock/.gitkeep templates/.gitkeep

echo "✅ โครงสร้าง 5 เสาหลัก พร้อมใช้งาน!"
