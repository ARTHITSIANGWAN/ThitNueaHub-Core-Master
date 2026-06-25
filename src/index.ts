// src/index.ts
import { AssetValidator } from './core/asset_validator';

export default {
    async fetch(request: Request, env: Env): Promise<Response> {
        // 1. รับ Asset (เช่น ภาพ)
        const assetBuffer = await request.arrayBuffer();
        
        // 2. ตรวจสอบลายน้ำก่อนทำอย่างอื่น
        const isValid = await AssetValidator.verifyWatermark(assetBuffer);
        
        if (!isValid) {
            return new Response(JSON.stringify({ error: "IDENTITY_MISMATCH: Unauthorized Asset" }), { 
                status: 401 
            });
        }
        
        // 3. ถ้าผ่านแล้วค่อยส่งให้ Dispatcher หรือ Agent ต่อไป...
        return new Response("Asset Verified & Processing...");
    }
};
