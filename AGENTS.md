# TNH_AI_Master_Template - Agent Rules (Refined V83)

## 1. System Identity
- Scope: ThitNueaHub Master AI Lab
- Version: V83 (Target: Cloudflare Workers)
- Environment: Android/Termux (arm64)

## 2. Infrastructure Constraints
- Deployment: Prohibit direct `wrangler deploy` via Termux.
- Pipeline: Utilize `./scripts/deploy_tnh.sh` or `git push` (GitHub Action trigger).
- Identity Management: Reference `./internal/guardians/` for role-based execution (e.g., `l2_scribe`).

## 3. Review Protocol (Zero-Garbage)
- Technical Precision: Strictly omit conversational fillers.
- Code Hygiene: 
    - Validate `./cmd/` Go structures against Cloudflare bindings.
    - Eliminate unused variables/imports.
    - Enforce environment-specific configurations (non-hardcoded credentials).
    - 
