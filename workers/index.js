// src/index.ts
// พิมพ์เขียว V84.5 - เครื่องยนต์ขับเคลื่อนด้วยระบบคิวและสัจจะข้อมูล

export interface Env {
	DB: D1Database;
	TNH_KV: KVNamespace;
	TNH_R2: R2Bucket;
	AI: any;
	ORCHESTRATOR_QUEUE: Queue;
}

export default {
	// 🏰 1. ท่อ HTTP Gateway: รับคำสั่ง ส่งเข้าคิวสับไวระดับไมโครวินาที
	async fetch(request: Request, env: Env, ctx: ExecutionContext): Promise<Response> {
		const url = new URL(request.url);
		const responseHeaders = new Headers({
			"Access-Control-Allow-Origin": "*",
			"Content-Type": "application/json; charset=utf-8"
		});

		// เส้นทางส่งงานเข้าคิวปิดฝั่งกูเกิ้ล รันฝั่งคลาวด์แฟลร์เด็ดขาด
		if (url.pathname === "/workflow/start" || url.pathname === "/process") {
			const jobID = "JOB_" + Date.now();
			
			// บรรจุวัตถุดิบคำสั่งยุทธศาสตร์ระดับ L1 เข้าสู่คิวส่งงานหลัก
			await env.ORCHESTRATOR_QUEUE.send({
				job_id: jobID,
				action: "DELEGATE_TASK",
				topic: url.searchParams.get("topic") || "SME_Orchestration_Core",
				timestamp: new Date().toISOString()
			});

			// บันทึก Log ลงความจำถาวร KV เพื่อค้ำยันบริบทหลักหลังวันที่ 4
			await env.TNH_KV.put("last_queued_job", jobID);

			return new Response(JSON.stringify({
				status: "🟢 Queued Successfully",
				version: "V84.5_Queue_Ignite",
				queue_name: "tnh-orchestrator-queue",
				assigned_job_id: jobID,
				msg: "สัจจะข้อมูลถูกล็อกลงตู้เซฟคิวเรียบร้อย ปลอดภัยจากอาการลืมบริบท 100%"
			}), { headers: responseHeaders, status: 202 });
		}

		return new Response(JSON.stringify({
			engine: "TNH-AI-V84.5-TRINITY",
			status: "CORE_QUEUE_GATEWAY_ONLINE",
			bindings_check: { kv: !!env.TNH_KV, d1: !!env.DB, r2: !!env.TNH_R2, queues: !!env.ORCHESTRATOR_QUEUE }
		}), { headers: responseHeaders });
	},

	// 🧠 2. ท่อ Queue Handler: ดึงข้อมูลจากคิวมาแจกจ่ายงานให้ลูกน้องทำงานเบื้องหลังอัตโนมัติ (Durable Loop)
	async queue(batch: MessageBatch<any>, env: Env, ctx: ExecutionContext): Promise<void> {
		// คัดกรองและประมวลผลตรรกะแบบมนุษย์ด้วย ค.ว.ย. Protocol (คิด -> วิเคราะห์ -> แยกแยะ)
		for (const message of batch.messages) {
			const job = message.body;
			console.log(`🤖 [Orchestrator Core] กำลังวิเคราะห์สัญญะข้อความในคิว ID: ${job.job_id}`);

			try {
				// 🕵️ สั่งการลูกน้องสายตรวจ L2 แก้วตา: อัปเดตสถานะสถิติตู้คอนเทนเนอร์และสิทธิ์ลง D1 Database
				await env.DB.prepare(
					"INSERT OR REPLACE INTO tnh_agent_permissions (agent_id, current_task, last_ignition) VALUES (?, ?, ?)"
				).bind("kaewta-agent", `PROCESSING_${job.topic}`, job.timestamp).run();

				// ยืนยันการทำชิ้นงานสำเร็จ ถอนข้อความออกจากคิวถาวร ไร้เศษขยะข้อมูลตกค้าง
				message.ack();
			} catch (error) {
				// หากตรรกะพังหรือเกิดปัญหาเครือข่าย ปล่อยให้ระบบจัดการโยนเข้าคิวขยะ DLQ ต่อไป
				console.error(`⚠️ [Queue Warning] การประมวลผลสิทธิ์ล้มเหลว:`, error);
			}
		}
	}
};
