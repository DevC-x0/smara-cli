# Docs Gap Report

Report ini membandingkan source code Smara dengan struktur `docs-site/` saat ini. Tujuannya menjaga dokumentasi tetap mengikuti fitur aktual tanpa perlu install Understand Anything dulu.

## Status umum

Docs sekarang sudah mencakup:

- Getting started: installation, quickstart, configuration.
- Core concepts: ask mode, agent mode, memory, skills, tools, MCP, workflows.
- Guides: SSH/VPS, create skill, create workflow, deploy project, analyze codebase, knowledge graph, generate docs, provider setup, release process, web interface, platform bots, browser automation, multimodal tools, scheduled automation, workspaces, cloud memory.
- Reference: CLI commands, config, tool list, Web API, Graphify, skill format, analytics, sharing, policy, run history, update, troubleshooting.
- Examples: docs generation, docs-site maintenance, provider eval report, Browser E2E, React deploy, Go refactor, server monitoring, common workflows.
- Tooling: `scripts/audit-docs-cli.mjs` untuk audit coverage command vs docs.

## Source areas yang sudah terwakili

| Source area | Docs target | Status |
|---|---|---|
| `cmd/smara/*.go` | `reference/cli-commands.md`, `scripts/audit-docs-cli.mjs` | Covered baseline + audit |
| `internal/agent/builtin_tools.go` dan tools terkait | `reference/tool-list.md`, `core-concepts/tools.md` | Covered baseline |
| `internal/memory`, `cmd/smara/memory*.go` | `core-concepts/memory.md`, `guides/cloud-memory.md` | Covered deeper |
| `internal/skill`, `cmd/smara/skill*.go` | `core-concepts/skills.md`, `reference/skill-format.md` | Covered baseline |
| `internal/mcp`, `cmd/smara/mcp.go` | `core-concepts/mcp.md` | Covered baseline |
| `cmd/smara/ssh.go`, deploy command | `guides/use-vps-ssh.md`, `guides/deploy-project.md` | Covered baseline |
| `cmd/smara/graphify.go`, graphify internals | `guides/knowledge-graph.md`, `reference/graphify.md`, `guides/analyze-codebase.md` | Covered deeper |
| `cmd/smara/provider.go`, `cmd/smara/login.go`, `internal/llm` | `guides/provider-setup.md`, `reference/config.md`, `examples/provider-eval-report.md` | Covered deeper |
| `internal/config` | `reference/config.md`, `getting-started/configuration.md` | Covered baseline |
| `internal/web`, `web/src` | `guides/web-interface.md`, `reference/web-api.md` | Covered baseline+ |
| `cmd/smara/update.go`, release assets | `guides/release-process.md`, `reference/update.md` | Covered baseline+ |

## Gaps yang sudah ditutup

| Gap | Halaman/tooling | Status |
|---|---|---|
| Bot/platform serve | `guides/platform-bots.md` | Covered |
| Evaluation provider | `guides/evaluate-provider.md`, `examples/provider-eval-report.md` | Covered |
| Analytics | `reference/analytics.md` | Covered |
| Sharing metadata | `reference/sharing.md` | Covered |
| Web interface | `guides/web-interface.md` | Covered |
| Web API routes | `reference/web-api.md` | Covered baseline |
| Update command | `reference/update.md` | Covered |
| Release process | `guides/release-process.md` | Covered |
| Browser automation | `guides/browser-automation.md`, `examples/browser-e2e.md` | Covered |
| Workflow orchestration | `core-concepts/workflows.md`, `guides/create-workflow.md` | Covered |
| Multimodal tools | `guides/multimodal-tools.md` | Covered |
| Policy system | `reference/policy.md` | Covered |
| Run history/replay | `reference/run-history.md` | Covered |
| Scheduled automation | `guides/scheduled-automation.md` | Covered |
| Workspace/exploration | `guides/workspaces.md` | Covered |
| Graphify detailed reference | `reference/graphify.md` | Covered |
| Provider/custom model setup | `guides/provider-setup.md` | Covered |
| Cloud Memory deeper guide | `guides/cloud-memory.md` | Covered |
| Troubleshooting berbasis error nyata | `reference/troubleshooting.md` | Covered baseline+ |
| Homepage graph visual polish | `index.md`, `custom.css` | Covered static interactive links |
| Docs maintenance example | `examples/docs-site-maintenance.md` | Covered |
| Audit otomatis docs vs CLI | `scripts/audit-docs-cli.mjs` | Covered baseline |

## Gaps tersisa / prioritas berikutnya

### 1. Interactive graph component — covered

Docs target:

```text
.vitepress/theme/components/DocsGraph.vue
index.md atau guides/knowledge-graph.md
```

Konten:

- graph docs clickable;
- import graph JSON kecil;
- filter kategori;
- tidak perlu library berat di tahap awal.

### 2. Web API detail per payload — covered baseline

`reference/web-api.md` sudah ada baseline endpoint list. Tahap lanjutan bisa menambahkan payload dan response aktual untuk endpoint prioritas:

```text
/api/chat
/api/web-sessions
/api/memories
/api/graph/query
/api/skills/run
```

### 3. More examples berbasis task nyata — covered baseline

Candidates:

```text
examples/vps-monitoring-skill.md
examples/release-checklist.md
examples/cloud-memory-sync.md
```

Konten:

- prompt siap pakai;
- expected output;
- safety checklist;
- kapan dibuat menjadi skill.

### 4. CI integration untuk docs audit — covered

`node scripts/audit-docs-cli.mjs` sudah bisa dijalankan manual. Tahap berikutnya bisa menambahkan:

```text
.github/workflows/docs-audit.yml
npm script docs:audit
```

## Rekomendasi eksekusi berikutnya

Urutan terbaik:

1. Interactive docs graph component ringan. — Covered
2. Perinci payload Web API untuk endpoint prioritas. — Covered baseline
3. Tambah example VPS monitoring skill dan release checklist. — Covered
4. Tambahkan `docs:audit` script atau CI workflow. — Covered

## Rule keselamatan

Saat memperbarui docs, jangan overwrite file runtime user kecuali user eksplisit meminta:

```text
internal/web/multisession.go
internal/web/session_handlers.go
web/src/api.ts
web/src/pages/Chat.tsx
```

Jika file-file itu muncul di `git status`, anggap sebagai baseline dari pekerjaan user sebelumnya.
