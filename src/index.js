// ตัวอย่างลอจิกใน workers/index.js (Cloudflare)
export default {
  async fetch(request, env) {
    // อ่านค่าสวิตช์จาก Environment Variable (เปลี่ยนค่าได้จาก CI/CD หรือ Dashboard)
    const activeBackend = env.ACTIVE_BACKEND; // ค่าอาจจะเป็น "CLOUDFLARE_ONLY", "GCP_CLOUD_RUN", หรือ "SPLIT"

    if (activeBackend === "GCP_CLOUD_RUN") {
        // ถ้าเปิดสวิตช์ GCP ไว้ ให้โยน Request ไปหา Go MCP Server ที่ Google Cloud Run
        const gcpUrl = "https://[YOUR-CLOUD-RUN-URL]/v1/mcp/ingress";
        return fetch(gcpUrl, request);
    }

    // ถ้าปิด GCP อยู่ ให้ Cloudflare ตอบกลับเอง หรือใช้ฟังก์ชันสำรอง
    return new Response(JSON.stringify({
        status: "SUCCESS",
        message: "Handled by Cloudflare Edge (GCP is currently sleeping)"
    }), { headers: { "Content-Type": "application/json" } });
  }
}
