#!/data/data/com.termux/files/usr/bin/bash
# PROJECT NAME: TNH-AI-MASTER-TEMPLATE
# FOLDER PATH: /scripts/
# FILE NAME: health_check.sh

# Project Name: tnh_ai
# Folder Path: /scripts/
# File Name: health_check.sh

set -e

echo "=== [1/4] 5S SSOT Directory Check ==="
if [ -d "$HOME/TNH-AI-MASTER-TEMPLATE" ] && [ ! -L "$HOME/TNH-AI-MASTER-TEMPLATE" ]; then
  mkdir -p "$HOME/archive_legacy"
  mv "$HOME/TNH-AI-MASTER-TEMPLATE" "$HOME/archive_legacy/backup_$(date +%Y%m%d_%H%M%S)"
  ln -s "$HOME/projects/TNH-AI-MASTER-TEMPLATE" "$HOME/TNH-AI-MASTER-TEMPLATE"
  echo "✅ Unified duplicate folder into ~/projects/TNH-AI-MASTER-TEMPLATE"
else
  echo "✅ Single Source of Truth verified"
fi

echo -e "\n=== [2/4] Git SSH & Security Guard ==="
cd "$HOME/projects/TNH-AI-MASTER-TEMPLATE"
git remote set-url origin git@github.com:ARTHITSIANGWAN/TNH-AI-MASTER-TEMPLATE.git
grep -q "^\.env$" .gitignore || printf "\n.env\ntnh-core\n" >> .gitignore
echo "✅ Git Remote locked to SSH & .gitignore verified"

echo -e "\n=== [3/4] Pure Go Core (:2026) & Latency Check ==="
if ! curl -s http://localhost:2026/health > /dev/null; then
  go build -ldflags="-s -w" -o tnh-core ./cmd/orchestra/main.go
  tmux kill-session -t tnh_core 2>/dev/null || true
  tmux new -d -s tnh_core "cd $HOME/projects/TNH-AI-MASTER-TEMPLATE && ./tnh-core"
  sleep 1
fi
curl -w "\nLatency: %{time_total}s\n" -s http://localhost:2026/health

echo -e "\n=== [4/4] Git Status & Hardware Node ==="
git status -s
echo "Free RAM: $(free -h | awk '/^Mem:/ {print $4 "/" $2}')"
