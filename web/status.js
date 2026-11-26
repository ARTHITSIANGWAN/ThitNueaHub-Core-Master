// Project Name: TNH-AI-MASTER-TEMPLATE
// Folder Path: /web/
// File Name: status.js

async function updateSystemStatus() {
  const statusEl = document.getElementById('live-telemetry');
  if (!statusEl) return;
  try {
    const start = performance.now();
    const res = await fetch('/health');
    const data = await res.json();
    const latency = (performance.now() - start).toFixed(2);
    statusEl.textContent = `โหนด: ${data.node} | RAM: ${data.go_alloc_mb} | ว่าง: ${data.free_disk_gb} | ตอบสนอง: ${latency}ms`;
  } catch (e) {
    statusEl.textContent = 'โหมดออฟไลน์: รอการเชื่อมต่อ Pure Go Core (:2026)';
  }
}

setInterval(updateSystemStatus, 5000);
updateSystemStatus();
