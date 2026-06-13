#!/bin/bash
# ⚔️ สคริปต์สั่งการกองทัพมหาอำนาจ ThitNueaHub V8.3
# แก้ไขใส่ชื่อโฟลเดอร์จริงหลังบ้านตรงล็อก เพื่อไม่ให้ท่อหลอนพ่น Error
PROJECTS=(
    "TNH_AI_JUST_A_TINY_AGENT"        # ขุนพลสมองกลหลัก (ลิปโป้ตัวตึงของพี่!)
    "tnh_ai_fortress_android"         # กล่องท่อแอนดรอยด์คุมเครื่อง OPPO
    "TNH_UI_UX_Project_Template"      # หน้ากาก UI/UX ส่วนสาธารณะที่เพิ่งแยกค่าย
    "tnh_sme_retail"                  # สล็อตลูกค่ายค้าปลีกย่อย
    "tnh_b2b_wholesale"               # สล็อตลูกค่ายค้าส่ง B2B
)

echo "⚔️ เริ่มการสั่งการกองทัพ (Deploying to Cloudflare)..."

for proj in "${PROJECTS[@]}"; do
    echo ">> สั่งการระบบไปยังมอดูล: $proj"
    cd "../../$proj" || continue # ถอยท่อกลับไปหาโฟลเดอร์หลักให้ถูกระดับ
    
    # คำสั่งลั่นไก Deploy ขึ้นคลาวด์แฟลร์แจกจ่าย DNA
    npx wrangler deploy --name "$proj-worker"
    
    cd - > /dev/null
done

echo "✅ ภารกิจเสร็จสิ้น: กองทัพ 11 ขุนพลออนไลน์หนึบแน่นหนา 2000%!"
