# PROJECT NAME: TNH-AI-THITNUEAHUB
# FOLDER PATH: /scripts/devops/
# FILE NAME: setup_gemini_mcp.sh

#!/data/data/com.termux/files/usr/bin/bash
set -euo pipefail

mkdir -p "$HOME/.gemini"

cat << 'JSON_EOF' > "$HOME/.gemini/settings.json"
{
  "mcpServers": {
    "github": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-github"],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "$GITHUB_PERSONAL_ACCESS_TOKEN"
      }
    },
    "devpost": {
      "httpUrl": "https://devpost.com/mcp"
    }
  }
}
JSON_EOF

jq . "$HOME/.gemini/settings.json" >/dev/null
echo "✅ สร้าง ~/.gemini/settings.json สำเร็จและโครงสร้าง JSON ถูกต้อง 100%"
