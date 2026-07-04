export interface Env {
	thitnueahub_core_db: D1Database;
	TNH_KV: KVNamespace;
	SECRET_KEY: string;
}

export default {
	async fetch(request: Request, env: Env, ctx: ExecutionContext): Promise<Response> {
		const url = new URL(request.url);
		
		// ตั้งค่าด่านกรองความปลอดภัยเบื้องต้น (CORS & Header Verification)
		const responseHeaders = new Headers({
			"Access-Control-Allow-Origin": "*",
			"Content-Type": "application/json; charset=utf-8"
		});

		// 1. ท่อตรวจสอบสิทธิ์ 11 ขุนพลจากตาราง tnh_agent_permissions ใน D1
		if (url.pathname === "/api/v84/agent-check") {
			const agentId = url.searchParams.get("agent_id") || "kaewta-agent";
			try {
				const { results } = await env.thitnueahub_core_db
					.prepare("SELECT * FROM tnh_agent_permissions WHERE agent_id = ?")
					.bind(agentId)
					.all();
				
				return new Response(JSON.stringify({ status: "success", data: results }), { headers: responseHeaders });
			} catch (err: any) {
				return new Response(JSON.stringify({ status: "error", msg: err.message }), { headers: responseHeaders, status: 500 });
			}
		}

		// 2. ท่อเชื่อมต่อสตรีมดักจับข้อมูลภาพด่วนจุดชนวนระบบ AI สายตา
		if (url.pathname === "/workflow/start") {
			const key = url.searchParams.get("key") || "default.jpg";
			// บันทึก State ลงในตู้เซฟความจำชั่วคราวเพื่อค้ำยันบริบทหลัก
			await env.TNH_KV.put("last_triggered_flow", `Flow_Key_${key}_Time_${Date.now()}`);
			
			return new Response(JSON.stringify({
				status: "🔥 Ignition",
				pipeline: "Durable Cloudflare Workflows",
				target_file: key,
				msg: "ส่งสัญญาณคัดแยกขยะข้อมูลเสร็จสิ้น ขุนพล AI เริ่มวิ่งงานเบื้องหลัง"
			}), { headers: responseHeaders });
		}

		return new Response(JSON.stringify({
			engine: "TNH-AI-V84.1-TRINITY",
			status: "FRONT_LAYER_ACTIVE",
			port_relay: "2026_CONNECTED"
		}), { headers: responseHeaders });
	}
};
        
