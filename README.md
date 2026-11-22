<!--
PROJECT NAME: ThitNueaHub-Core-Master
FOLDER PATH: /
FILE NAME: README.md
-->

# ThitNueaHub-Core-Master (`TNH-AI-THITNUEAHUB`)

> Sovereign Pure Go Backend Infrastructure, 11-Layer Multi-Agent Governance, and Mobile-First Edge Orchestration (`#tnhAi` `#thitnueahub`).

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Storage](https://img.shields.io/badge/SQLite-WAL%20Mode-003B57?logo=sqlite&logoColor=white)](https://www.sqlite.org/wal.html)
[![Runtime](https://img.shields.io/badge/Runtime-Android%20Termux%20%26%20Linux-2496ED)](https://termux.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

---

## Overview

**ThitNueaHub-Core-Master** is the foundational architecture repository for the **TNH-AI-THITNUEAHUB** ecosystem. Built with pure Go and zero external framework bloat, it is engineered for deterministic execution on resource-constrained mobile environments (Android Termux) as well as standard Linux containers.

### Core Engineering Principles
* **Zero-Garbage Discipline**: Minimal allocations, explicit error handling, and clean file boundaries with mandatory project metadata headers.
* **Strict 11-Layer Separation (`L1`–`L11`)**: Every agent and service module operates within a single, non-overlapping responsibility boundary.
* **Edge-First Privacy**: Customer payloads are processed at the endpoint/client boundary without unauthorized central retention; only legally required insight records are persisted.
* **Cryptographic Perimeter**: Inter-layer envelopes are authenticated using environment-injected `HMAC-SHA256` signatures.

---

## System Architecture

```mermaid
graph TD
    Client[Client / Edge Endpoint] --> L8[L8 Glock: HMAC-SHA256 Guard]
    L8 --> L4[L4 Bridge: API & Android Gateway]
    L4 --> L7[L7 Oracle: Reasoning & Analytics]
    L4 --> L5[L5 Engine: Go Workers & SQLite WAL]
    L9[L9 Sentinel: Anomaly & Rate Guard] -.-> L4
    L10[L10 Compliance: Privacy & Audit] -.-> L5
    L11[L11 Balancer: Termux Resource Control] -.-> L5
