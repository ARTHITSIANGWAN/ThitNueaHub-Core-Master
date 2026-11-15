// PROJECT NAME: TNH-AI-MASTER-TEMPLATE
// FOLDER PATH: /cmd/
// FILE NAME: core_poller.ts

import { StrictValidator } from './validator';

async function runTNHLoop(env: Env) {
    try {
        // 1. ดึงข้อมูลผ่าน HTTPS (พร้อม Timeout)
        const rawOrders = await fetch('https://worker.yourdomain.workers.dev/getChat', { signal: AbortSignal.timeout(5000) });
        const orders = await rawOrders.json();

        // 2. Validate ก่อนส่งให้ Dispatcher (ป้องกัน Garbage)
        const validatedJob = StrictValidator.validate(orders);

        // 3. ส่งเข้า Dispatcher (เพื่อให้ 11 ขุนพลรับไปจัดการ)
        const result = await dispatchToAgent(env, validatedJob);
        
        // 4. ส่งกลับอย่างปลอดภัย
        await postToMessenger(result);

    } catch (err) {
        // บันทึกลง Log แทนการปล่อยให้ Error ทำงาน
        console.error("TNH_CORE_LOOP_ERROR:", err);
    }
}
