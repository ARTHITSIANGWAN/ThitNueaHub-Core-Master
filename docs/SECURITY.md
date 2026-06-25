# 🛡️ TNH_AI SECURITY POLICY

## 1. Core Principles
ระบบนี้ยึดถือความโปร่งใส (Transparency) และความปลอดภัย (Security) เป็นรากฐานสำคัญ ทุกโปรเจกต์ในเครือข่าย TNH_AI ต้องปฏิบัติตามกฎนี้โดยไม่มีข้อยกเว้น

## 2. Integrity & Governance
- **Auditing:** ทุกโปรเจกต์ต้องรัน `FiveTwo.ml` เพื่อตรวจสอบความสมบูรณ์ของ Code/Data อยู่เสมอ
- **Logging:** การเปลี่ยนแปลงที่สำคัญต้องถูกบันทึกใน `Chronicle.ml` เพื่อการตรวจสอบย้อนหลัง (Audit Trail)
- **Zero-Garbage:** ข้อมูลที่หมดอายุหรือเป็นช่องโหว่ ต้องถูกกำจัดทิ้งทันทีตามรอบ 5ส

## 3. Access Control
- **Environment:** ห้าม Hardcode API Keys หรือ Credentials ใดๆ ลงใน Code โดยเด็ดขาด ให้ใช้ `.env` และจัดการผ่านระบบที่ปลอดภัยเท่านั้น
- **Principle of Least Privilege:** ให้สิทธิ์เข้าถึงเฉพาะสิ่งที่จำเป็นต้องใช้ในการรันงานเท่านั้น

## 4. Vulnerability Reporting
หากพบช่องโหว่หรือความผิดปกติในระบบ:
1. แจ้งเหตุไปยัง Hub (Master) ทันที
2. บันทึกเหตุการณ์ลงใน `Chronicle.ml`
3. กักกัน (Isolate) ส่วนที่เสียหายจนกว่าการตรวจสอบจะเสร็จสิ้น

## 5. Scope
นโยบายนี้มีผลบังคับใช้กับทุกโปรเจกต์ที่เชื่อมต่อกับ TNH_AI_MASTER (Hub) และถือเป็นส่วนหนึ่งของ DNA มาตรฐาน

---
*Last Updated: 2026-05-27*
*Governance by: TNH_AI_MASTER Auditor (FiveTwo.ml)*
