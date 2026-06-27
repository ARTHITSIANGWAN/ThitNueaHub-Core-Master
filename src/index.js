export default {
  async fetch(request, env, ctx) {
    const url = new URL(request.url);

    // ==========================================
    // 🎛️ 1. สวิตช์ควบคุม (Feature Toggle)
    // ==========================================
    // ดึงค่ามาจาก Environment Variable ใน Cloudflare
    const useCloudRun = env.USE_CLOUD_RUN === "true";

    // ==========================================
    // 🔵 2. กรณีสวิตช์ "เปิด" (โยนงานหนักไป Google Cloud Run)
    // ==========================================
    if (useCloudRun && url.pathname === "/dispatch") {
      const cloudRunUrl = env.CLOUD_RUN_URL; // เช่น https://fortress-xxx-as.a.run.app
      
      // คัดลอก Request เดิมทั้งหมด แล้วเปลี่ยนปลายทางไปที่ Cloud Run
      const proxyRequest = new Request(`${cloudRunUrl}/dispatch`, {
        method: request.method,
        headers: request.headers,
        body: request.body
      });
      
      // โยนงานไปแล้วรอรับผลลัพธ์กลับมาส่งให้ User
      return fetch(proxyRequest);
    }

    // ==========================================
    // 🟢 3. กรณีสวิตช์ "ปิด" (Cloudflare ตอบกลับเองทันที)
    // ==========================================
    // ใช้สำหรับรับแขกทั่วไป หรือตอนที่เราปิดเซิร์ฟเวอร์ Cloud Run เพื่อประหยัดเงิน
    const edgeResponse = {
      status: "success",
      source: "Cloudflare Edge (Front-door)",
      message: "🚀 ระบบศูนย์กลาง TNH ทำงานปกติ (โหมดประหยัดพลังงาน)",
      timestamp: new Date().toISOString(),
      path_requested: url.pathname
    };

    return new Response(JSON.stringify(edgeResponse, null, 2), {
      status: 200,
      headers: { 
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*" // ป้องกันปัญหา CORS
      }
    });
  },
};
