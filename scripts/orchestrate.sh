#!/bin/bash
# สคริปต์สั่งการ 11 ขุนพล
PROJECTS=("Project_01" "Project_02" "Project_03" "Project_04" "Project_05" \
          "Project_06" "Project_07" "Project_08" "Project_09" "Project_10" "Project_11")

echo "⚔️ เริ่มการสั่งการกองทัพ (Deploying to Cloudflare)..."

for proj in "${PROJECTS[@]}"; do
    echo ">> สั่งการไปยัง: $proj"
    cd "../$proj" || continue
    
    # คำสั่งเชื่อมต่อ Cloudflare (ถ้าลง Wrangler ไว้)
    npx wrangler deploy --name "$proj-worker"
    
    cd - > /dev/null
done

echo "✅ ภารกิจเสร็จสิ้น: 11 ขุนพลออนไลน์แล้ว!"
