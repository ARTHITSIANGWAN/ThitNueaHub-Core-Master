# 🏰 TNH_AI_Master_Template

แม่แบบมาตรฐานสำหรับโครงการ TNH_AI ทั้ง 12 แห่ง

## 🚀 The Hub Logic
- **Master Template:** แหล่งรวม DNA ต้นฉบับ
- **Deployment:** ขับเคลื่อนด้วย Cloudflare Workers เพื่อการรันงานระดับ Edge
- **Policy:** ควบคุมผ่าน SYNC_POLICY.md และจัดการสถานะด้วย Chronicle.ml

## 🛠️ Infrastructure Overview
โครงสร้างถูกแบ่งเป็น 5 เสาหลัก:
1. **Docs:** คู่มือและกฎระเบียบ
2. **Services:** งานเฉพาะทาง (Scraper/Scheduler)
3. **Data:** ที่จัดเก็บผลลัพธ์และคลังคอนเทนต์
4. **Templates:** DNA ต้นฉบับ
5. **Scripts:** เครื่องมืออัตโนมัติ (Automation)

## ⚖️ Governance
การเปลี่ยนแปลงใดๆ ใน Template ต้องได้รับอนุมัติผ่าน SYNC_POLICY.md และบันทึกโดย Chronicle.ml เสมอ
