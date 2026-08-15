# Analyze Codebase

Smara doc site memakai konsep **Understand Anything = analisis struktur + knowledge graph**, lalu hasilnya dijadikan bahan dokumentasi VitePress.

Untuk tahap awal, analisis bisa dilakukan langsung oleh Smara tanpa meng-install Understand Anything eksternal: Smara membaca source code, grep command, membuat graph/call graph bila perlu, lalu menulis ulang Markdown docs.

## Target analisis repo Smara

| Area | Sumber utama | Output docs |
|---|---|---|
| CLI command | `cmd/smara/*.go` | `reference/cli-commands.md` |
| Agent mode & safety | `internal/agent`, `pkg/agent`, `internal/safety` | `core-concepts/agent-mode.md`, `core-concepts/tools.md` |
| Memory | `cmd/smara/memory*.go`, `internal/memory` | `core-concepts/memory.md`, `guides/cloud-memory.md` |
| Skills | `cmd/smara/skill*.go`, `internal/skill`, `skills/*.json` | `core-concepts/skills.md`, `reference/skill-format.md` |
| MCP | `cmd/smara/mcp.go`, `internal/mcp`, `pkg/mcp` | `core-concepts/mcp.md` |
| SSH/VPS | `cmd/smara/ssh.go`, `internal/ssh` | `guides/use-vps-ssh.md` |
| Graphify | `cmd/smara/graphify.go`, `internal/graphify`, `graphify-out/` | `guides/knowledge-graph.md` |
| Web UI | `internal/web`, `web/src` | guide/reference web docs bila diperlukan |

## Workflow

<div class="smara-flow">
  <div class="smara-card"><strong>Scan</strong><p>Baca struktur repo, package, command, handler, config, README, dan dokumentasi lama.</p></div>
  <div class="smara-card"><strong>Cluster</strong><p>Kelompokkan fitur menjadi getting started, concepts, guides, reference, dan examples.</p></div>
  <div class="smara-card"><strong>Map</strong><p>Buat peta relasi antara fitur: agent, skills, memory, MCP, SSH, tools, web, dan graphify.</p></div>
  <div class="smara-card"><strong>Draft</strong><p>Generate Markdown awal berdasarkan source code dan contoh penggunaan nyata.</p></div>
  <div class="smara-card"><strong>Verify</strong><p>Build docs, cek link, dan review konten manual.</p></div>
</div>

## Command analisis praktis

Cari command Cobra:

```bash
grep -R "cobra.Command" -n cmd/smara
```

Lihat command yang didaftarkan ke root:

```bash
grep -R "rootCmd.AddCommand" -n cmd/smara
```

Bangun knowledge graph Go:

```bash
smara graphify init . --name smara
smara graphify query "memory cloud and graphify docs" --name smara --depth 2
```

Build docs setelah update:

```bash
cd docs-site
npm run docs:build
```

## Prompt contoh

```text
analisis repo ini untuk membuat dokumentasi: identifikasi fitur utama, command, config, dan gap docs
```

```text
buat peta fitur Smara dari source code lalu rekomendasikan struktur VitePress docs
```

```text
update reference/cli-commands.md berdasarkan cmd/smara/*.go tanpa menyentuh file bug fix web
```

## Output yang dicari

- daftar fitur utama aktual
- daftar command CLI aktual
- daftar tool/skill penting
- flow onboarding user baru
- gap antara source code, README, dan docs-site
- draft halaman Markdown

## Hubungan dengan VitePress

VitePress tetap menjadi publishing layer. Analisis codebase hanya dipakai sebagai bahan untuk menulis halaman docs yang lebih akurat.

```text
Analyze source -> Map feature graph -> Write Markdown -> Build VitePress -> Publish
```
