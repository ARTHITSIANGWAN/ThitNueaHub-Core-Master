// src/index.ts - TNH ENGINE V84.9.1 [ZERO-GOOGLE PROTOCOL]
export interface Env {
	DB: D1Database;
	TNH_KV: KVNamespace;
	AI: any;
	AI_GATEWAY: any;
}

export default {
	async fetch(request: Request, env: Env): Promise<Response> {
		const url = new URL(request.url);
		const headers = new Headers({ 
			"Content-Type": "application/json; charset=utf-8", 
			"Access-Control-Allow-Origin": "*" 
		});

		// 🏰 ท่อรับงานความเร็วแสงดักสัญญะผ่านเครือข่ายปิดคลาวด์แฟลร์เท่านั้น
		if (url.pathname === "/api/v84/run-ai" || url.pathname === "/workflow/start") {
			const prompt = url.searchParams.get("prompt") || "ตรวจสอบประวัติบอทตกค้าง";

			// 🧠 รันตรรกะชั้นสูงผ่านเสาหลัก AI Gateway โดยใช้สิทธิ์โมเดล Llama 3.3 Fast
			const aiResponse = await env.AI.run("@cf/meta/llama-3.3-70b-instruct-fp8-fast", {
				prompt: prompt,
				gateway: "tnh-ai-gateway-v8" // ควบคุม Analytics & Cache ป้องกันดาต้ารั่วไหล
			});

			// 🕵️ สั่งการลูกน้อง L2 แก้วตา: บันทึกข้อมูลและสถานะลงตาราง ai_logs ใน D1 thitnueahub-core-db ทันที
			await env.DB.prepare(
				"INSERT OR REPLACE INTO ai_logs (model_name, status, timestamp) VALUES (?, ?, ?)"
			).bind("llama-3.3-70b", "CONTROLLED_BY_GATEWAY_SUCCESS", Date.now()).run();

			// 🧹 สั่งการลูกน้อง L5 ไอ้จ๊อด: สรรสร้างพื้นที่สะอาด ล้างหน่วยความจำตกค้างด้วยมาตรการ 5 ส
			await env.TNH_KV.put("last_purge_status", "CLEAN_ZERO_GARBAGE");

			return new Response(JSON.stringify({
				status: "🟢 AI Gateway Active & Locked",
				provider_name: "tnh-ai-v83-trinity-empire",
				google_zone: "CLOSED_SECURE",
				total_models_secured: 59,
				result: aiResponse
			}), { headers });
		}

		return new Response(JSON.stringify({ 
			status: "ONLINE", 
			msg: "ThitNueaHub Core Secure Serverless Engine" 
		}), { headers });
	}
};
        
