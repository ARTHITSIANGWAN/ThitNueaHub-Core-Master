# Agent System Instructions & Context

## Project Identity
- Name: TNH_AI_Master_Template (TRINITY-EMPIRE)
- Core Version: V83 (Active Platform)
- Environment: Cloudflare Workers & Mobile Termux (Android arm64)

## Architecture Norms (Zero-Garbage Protocol)
1. DO NOT generate unnecessary boilerplate code or hallucinated explanations.
2. Maintain clean, high-performance execution using Go and Python.
3. Edge deployment target is strictly Cloudflare Workers via clean GitHub Webhooks syncing.

## Infrastructure Constraints
- Local ports active: :2026 (Client), :2027 (Bot Go), :2028 (Security/Token)
- Strict Rule: Do not attempt to run local `workerd` simulation on Android arm64. Always sync via Git push.
- 
