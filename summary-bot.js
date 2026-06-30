/**
 * summary-bot.js (Kaewta Version 8.2 - Zero Garbage)
 * ระบบ Supervisor L2 สำหรับตรวจสอบออเดอร์ในอาณาจักร ThitNueaHub
 */

const WORKER_URL = 'https://tnh.9thera.workers.dev/getChat';
const INTERVAL_MS = 5 * 60 * 1000; // รอบละ 5 นาที

async function run() {
    console.log('🤖 [แก้วตา] กำลังเริ่มรอบการตรวจสอบออเดอร์...');
    
    try {
        const response = await fetch(WORKER_URL, {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'User-Agent': 'Kaewta-Supervisor-L2'
            }
        });

        if (!response.ok) {
            throw new Error(`เซิร์ฟเวอร์ตอบกลับ: ${response.status}`);
        }

        const rawText = await response.text();
        
        // เช็กด่าน Cloudflare
        if (rawText.trim().startsWith('<') || rawText.includes('<!DOCTYPE')) {
            console.error('⚠️ [แก้วตา] ติดด่าน Cloudflare: บอสลืมตั้ง Bypass ใน Dashboard หรือเปล่าครับ?');
            return;
        }

        const orders = JSON.parse(rawText);
        
        if (Array.isArray(orders) && orders.length > 0) {
            console.log(`📦 [แก้วตา] พบ ${orders.length} ออเดอร์ใหม่! กำลังส่งให้ Llama ประมวลผล...`);
        } else {
            console.log('✅ [แก้วตา] ไม่พบงานค้าง ระบบเรียบร้อยดีค่ะ');
        }

    } catch (error) {
        console.error('⚠️ [แก้วตา] เกิดข้อผิดพลาด:', error.message);
    } finally {
        console.log('⏳ รอตรวจสอบอีกครั้งใน 5 นาที...');
    }
}

// รันทันทีและตั้งเวลาวนลูป
run();
setInterval(run, INTERVAL_MS);
