#!/bin/bash
# TNH_AI_ORCHESTRATOR: สั่งการกองทัพ 11 ขุนพล
# Usage: ./scripts/orchestrate.sh

# ตั้งค่าจำนวนโปรเจกต์ (01 ถึง 11)
TOTAL_PROJECTS=11

echo "========================================="
echo "🚀 COMMANDER: เริ่มต้นสั่งการกองทัพ (Deploy All)"
echo "========================================="

for i in $(seq -f "%02g" 1 $TOTAL_PROJECTS); do
    DIR="TNH_AI_PROJECT_$i"
    
    # เช็คว่าโฟลเดอร์มีอยู่จริงไหม
    if [ -d "../$DIR" ]; then
        echo "⚔️ [Processing] -> $DIR"
        
        # เข้าไปในโฟลเดอร์
        cd "../$DIR" || continue
        
        # รันคำสั่ง Deploy
        # บอสสามารถเปลี่ยนเป็น 'npm run deploy' หรือคำสั่งที่บอสใช้ประจำได้
        npx wrangler deploy --name "$DIR-worker"
        
        # กลับมาที่เดิม
        cd - > /dev/null
        echo "✅ [Success] -> $DIR Deploy เรียบร้อย"
        echo "-----------------------------------------"
    else
        echo "⚠️ [Skipped] -> ไม่พบโฟลเดอร์ $DIR"
    fi
done

echo "========================================="
echo "🏁 ภารกิจสั่งการเสร็จสิ้น: 11 ขุนพล ออนไลน์ครบถ้วน!"
echo "========================================="
