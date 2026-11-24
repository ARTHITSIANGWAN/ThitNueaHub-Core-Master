# PROJECT NAME: TNH-AI-THITNUEAHUB
# FOLDER PATH: /scripts/simulation/
# FILE NAME: disaster_sandbox_suite.sh

#!/data/data/com.termux/files/usr/bin/bash
set -euo pipefail

GATEWAY_URL="http://127.0.0.1:2026/dispatch"
TS=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

echo "🚨 [1/3] Simulating Fire & Perimeter Lock (V4 F16 -> V5 Gripen)..."
curl -s -X POST "$GATEWAY_URL" -H "Content-Type: application/json" -d "{
  \"sender\": \"L9_SENTINEL_F16\",
  \"receiver\": \"L5_GRIPEN_COMMANDER\",
  \"action\": \"EMERGENCY_FIRE_DISPATCH\",
  \"priority\": \"CRITICAL_P0\",
  \"timestamp\": \"${TS}\",
  \"payload\": {
    \"mode\": \"SANDBOX_REAL_SCENARIO\",
    \"incident_type\": \"STRUCTURAL_FIRE\",
    \"zone_control\": \"PERIMETER_LOCKED_500M\",
    \"dispatch_units\": [\"FIRE_TRUCK_01\", \"RESCUE_TEAM_ALPHA\"],
    \"evacuation_priority\": \"VULNERABLE_GROUPS_FIRST\",
    \"offline_sync\": \"shared_truth.db\"
  }
}" >/dev/null || true

echo "🌊 [2/3] Simulating Nepal GLOF / Rapid Water Rate-of-Change..."
curl -s -X POST "$GATEWAY_URL" -H "Content-Type: application/json" -d "{
  \"sender\": \"L4_WATER_SENSOR_TELEMETRY\",
  \"receiver\": \"L5_GRIPEN_COMMANDER\",
  \"action\": \"EARLY_WARNING_FLASH_FLOOD\",
  \"priority\": \"CRITICAL_P0\",
  \"timestamp\": \"${TS}\",
  \"payload\": {
    \"mode\": \"SANDBOX_REAL_SCENARIO\",
    \"incident_type\": \"FLASH_FLOOD_WALL_OF_WATER\",
    \"current_level_m\": 2.10,
    \"critical_threshold_m\": 3.50,
    \"delta_5min_m\": 0.45,
    \"trigger_reason\": \"RAPID_RATE_OF_CHANGE\",
    \"evacuation_priority\": \"VULNERABLE_GROUPS_FIRST\",
    \"offline_sync\": \"shared_truth.db\"
  }
}" >/dev/null || true

echo "🌪️ [3/3] Simulating Tak Valley Severe Wind Gust / Downburst..."
curl -s -X POST "$GATEWAY_URL" -H "Content-Type: application/json" -d "{
  \"sender\": \"L4_WEATHER_TELEMETRY\",
  \"receiver\": \"L5_GRIPEN_COMMANDER\",
  \"action\": \"EARLY_WARNING_WIND_GUST\",
  \"priority\": \"CRITICAL_P0\",
  \"timestamp\": \"${TS}\",
  \"payload\": {
    \"mode\": \"SANDBOX_REAL_SCENARIO\",
    \"incident_type\": \"SEVERE_DOWNBURST_WIND_GUST\",
    \"location\": \"TAK_VALLEY_CORRIDOR\",
    \"wind_gust_kmh\": 78.5,
    \"pressure_drop_hpa\": -3.2,
    \"automated_actions\": [\"EVACUATE_HIGH_RISK_ZONES\", \"GROUND_ALL_UAV_DRONES\"],
    \"offline_sync\": \"shared_truth.db\"
  }
}" >/dev/null || true

echo "✅ All 3 Sandbox Disaster Scenarios Dispatched!"
