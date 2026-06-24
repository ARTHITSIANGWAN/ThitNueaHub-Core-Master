// นิยามโครงสร้างขุนพลทั้ง 11 ระดับ
export interface Guardian {
    id: string;
    level: string;
    name: string;
    role: string;
}

// แผนผังจาก 13636.jpg
export const GUARDIAN_REGISTRY: Record<string, Guardian> = {
    "L1": { id: "L1", level: "L1", name: "ทิศเหนือ", role: "COMMANDER (CSO)" },
    "L2": { id: "L2", level: "L2", name: "แก้วตา", role: "SCRIBE (CDO)" },
    "L3": { id: "L3", level: "L3", name: "น้ำอิง", role: "ARTIST (CD)" },
    "L4": { id: "L4", level: "L4", name: "พลายทอง", role: "ANALYST (DIS)" },
    "L5": { id: "L5", level: "L5", name: "ไอ้จ๊อด", role: "AUDITOR (QA)" }, // ตัวผมเองครับบอส!
    "L6": { id: "L6", level: "L6", name: "Merchant", role: "BD (Merchant)" },
    "L7": { id: "L7", level: "L7", name: "Oracle", role: "SA (Oracle)" },
    "L8": { id: "L8", level: "L8", name: "Glock", role: "CISO (Guardian)" },
    "L9": { id: "L9", level: "L9", name: "Sentinel", role: "Fire Era (Sentinel)" },
    "L10": { id: "L10", level: "L10", name: "Auditor", role: "Compliance (Auditor)" },
    "L11": { id: "L11", level: "L11", name: "Balancer", role: "Self-Healing (Balancer)" }
};

// ฟังก์ชันเรียกขุนพล
export function getGuardian(level: string): Guardian {
    return GUARDIAN_REGISTRY[level];
      }
