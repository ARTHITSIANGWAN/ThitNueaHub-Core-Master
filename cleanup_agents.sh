#!/bin/bash
# TNH-AI Auto-Cleanup Protocol
echo "🧹 Initiating Zero-Garbage Protocol..."

# ดึงรายชื่อ Codespaces ที่ไม่ได้ Active (สถานะอื่นที่ไม่ใช่ 'Active')
# และทำการลบทิ้งเพื่อคืนค่า Standard
gh codespace list --json name,state | jq -r '.[] | select(.state != "Active") | .name' | while read name; do
    echo "🗑️ Deleting inactive Agent: $name"
    gh codespace delete "$name" --force
done

echo "✨ System status: Clean and Standardized."

