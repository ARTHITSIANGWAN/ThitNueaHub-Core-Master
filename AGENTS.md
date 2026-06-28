# TNH_AI_Master_Template - Agent Rules

## 1. System Identity & Environment
- **Project Scope:** ThitNueaHub Master AI Lab Infrastructure
- **Active Version:** V83 (Core Target Platform)
- **Host OS:** Android (via Termux Environment / arm64 architecture)

## 2. Infrastructure Guide & Constraints
- **Deployment Protocol:** Do not use local `wrangler deploy` in Termux directly due to arm64/workerd binary restrictions.
- **Sync Method:** Always run scripts under `./scripts/` (e.g., `deploy_tnh.sh`) or execute `git push` to let GitHub Webhooks trigger Cloudflare production deployment.
- **Internal Guardians:** Reference directory `./internal/guardians/` for specific agent identities (e.g., Use `l2_scribe` for content management and supervisor alignment).

## 3. Zero-Garbage Review Protocol
- Maximize technical precision. Avoid prose or repetitive conversational filler.
- Ensure all Go structures under `./cmd/` and configuration flags map directly to active Cloudflare bindings without generating unused variables.
- 
