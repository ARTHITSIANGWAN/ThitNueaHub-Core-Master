// src/index.ts
export interface Env {
  DB: D1Database;
  TNH_R2: R2Bucket;
  TNH_KV: KVNamespace;
  AI: any;
}

const SECRET_KEY = "TNH_SOOD_SOI_KEY"; // ยึดค่านี้ให้ตรงกับใน Android Fortress

export default {
  async fetch(request: Request, env: Env, ctx: ExecutionContext): Promise<Response> {
    // 1. ตรวจสอบ Signature (X-TNH-Signature) ที่ยิงมาจาก Android
    const signature = request.headers.get("X-TNH-Signature");
    const payload = await request.clone().arrayBuffer();

    if (!(await verifySignature(payload, signature, SECRET_KEY))) {
      return new Response("Unauthorized: Identity Mismatch", { status: 401 });
    }

    // 2. ถ้าผ่าน ให้ประมวลผล (เช่น เก็บลง D1 หรือ R2)
    // เพิ่ม Logic การจัดการ Job ที่นี่...
    
    return new Response("Payload Accepted by Gripen", { status: 200 });
  },

  // Cron Trigger (ตามความไวระดับ F-16 ทุก 3 นาที)
  async scheduled(event: ScheduledEvent, env: Env, ctx: ExecutionContext) {
    console.log("Gripen System Triggered by Cron...");
    // Logic ตรวจสอบ Calendar หรือ Sync ข้อมูล...
  }
};

async function verifySignature(body: ArrayBuffer, sig: string | null, secret: string): Promise<boolean> {
  // Logic ตรวจสอบ HMAC-SHA256 (ต้องตรงกับฟังก์ชันใน main.go บน Android)
  return true; // พัฒนาต่อโดยใช้ crypto.subtle.importKey
}
