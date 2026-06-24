// src/index.ts
import { StrictValidator } from './core/validator';
import { runDispatcher } from './core/dispatcher';

export default {
    async fetch(request: Request, env: Env): Promise<Response> {
        try {
            // 1. รับข้อมูลดิบ
            const rawPayload = await request.json();
            
            // 2. Validate (คัดกรอง Noise/Garbage ออก)
            const cleanPayload = StrictValidator.validate(rawPayload);
            
            // 3. Dispatch (ส่งไปให้ Agent ที่เกี่ยวข้อง)
            const result = await runDispatcher(env, cleanPayload);
            
            return new Response(JSON.stringify(result), { 
                status: 200, 
                headers: { 'Content-Type': 'application/json' } 
            });
            
        } catch (error: any) {
            // ถ้า Validate ไม่ผ่าน หรือเกิด Error จะถูกดักไว้ที่นี่
            return new Response(JSON.stringify({ error: error.message }), { 
                status: 403 
            });
        }
    }
};
              
