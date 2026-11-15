#!/bin/bash
# PROJECT NAME: TNH-AI-MASTER-TEMPLATE
# FOLDER PATH: /scripts/
# FILE NAME: purge_workflows.sh

# --- TNH Purge Protocol: กวาดล้าง Workflow ขยะให้สิ้นซาก ---
# กฎ: ลบทุกอย่างใน .github/workflows/ ยกเว้นที่ฐานบัญชาการใหญ่ (TNH_AI_Master_Template)

echo "⚠️ [Purge Protocol] เริ่มต้นการกวาดล้าง Workflow ในอาณาจักร..."

# รายชื่อโฟลเดอร์โปรเจกต์ (01-20)
PROJECTS=(TNH_AI_PROJECT_01 TNH_AI_PROJECT_02 TNH_AI_PROJECT_03 TNH_AI_PROJECT_04 TNH_AI_PROJECT_05 TNH_AI_PROJECT_06 TNH_AI_PROJECT_07 TNH_AI_PROJECT_08 TNH_AI_PROJECT_09 TNH_AI_PROJECT_10 TNH_AI_PROJECT_11 TNH_AI_PROJECT_12 TNH_AI_PROJECT_13 TNH_AI_PROJECT_14 TNH_AI_PROJECT_15 TNH_AI_PROJECT_16 TNH_AI_PROJECT_17 TNH_AI_PROJECT_18 TNH_AI_PROJECT_19 TNH_AI_PROJECT_20)

for PROJ in "${PROJECTS[@]}"; do
    if [ -d "$PROJ/.github/workflows" ]; then
        echo "🔥 กวาดล้าง Workflow ใน: $PROJ"
        rm -rf "$PROJ/.github/workflows"
    else
        echo "✅ $PROJ สะอาดอยู่แล้ว..."
    fi
done

echo "🏁 [Purge Complete] อาณาจักรสะอาดไร้ Workflow แปลกปลอมแล้วบอส!"

