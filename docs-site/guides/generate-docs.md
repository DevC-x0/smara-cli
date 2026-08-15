# Generate Docs

Smara dapat dipakai untuk membantu membuat dan memperbarui dokumentasi Markdown. Konsepnya menggabungkan analisis ala Understand Anything, publishing VitePress, dan refinement oleh LLM.

## Stack rekomendasi

```text
Understand Anything = analisis struktur + knowledge graph
VitePress = doc site utama
Smara/LLM = generate dan refine konten Markdown
```

Dalam praktik saat ini, Smara bisa langsung melakukan tahap analisis tanpa install tool eksternal terlebih dahulu. Understand Anything tetap berguna nanti sebagai audit/gap-analysis tambahan.

## Flow produksi docs

1. **Scan repo** — baca `README.md`, `SMARA.md`, `cmd/`, `internal/`, `pkg/`, `web/`, dan docs existing.
2. **Map fitur** — kelompokkan menjadi onboarding, konsep, guide, reference, dan examples.
3. **Extract reference** — ambil command dari `cmd/smara/*.go`, tool dari runtime/tool list, config dari `internal/config` dan README.
4. **Generate Markdown** — tulis halaman di `docs-site/**/*.md`.
5. **Refine** — polish bahasa, tambah contoh command, dan hilangkan klaim yang tidak ada di source.
6. **Verify** — build VitePress dan cek link.
7. **Publish** — deploy output `.vitepress/dist` ke Vercel/static hosting.

## Struktur docs

```text
docs-site/
  getting-started/
    installation.md
    quickstart.md
    configuration.md

  core-concepts/
    ask-mode.md
    agent-mode.md
    memory.md
    skills.md
    tools.md
    mcp.md

  guides/
    use-vps-ssh.md
    create-skill.md
    deploy-project.md
    analyze-codebase.md
    knowledge-graph.md
    generate-docs.md
    cloud-memory.md

  reference/
    cli-commands.md
    config.md
    tool-list.md
    skill-format.md
    troubleshooting.md

  examples/
    docs-generation.md
    react-deploy.md
    golang-refactor.md
    server-monitoring.md
```

## Prompt contoh untuk Smara

Generate feature guide:

```text
buatkan dokumentasi VitePress untuk fitur skill Smara berdasarkan source code cmd/smara/skill.go, internal/skill, dan README
```

Audit gap docs:

```text
cek gap docs: fitur apa di cmd/smara dan README yang belum ada di docs-site? jangan edit dulu, buat daftar prioritas
```

Update reference:

```text
update docs-site/reference/cli-commands.md berdasarkan cmd/smara/*.go, lalu build docs
```

Refine halaman:

```text
refine halaman Markdown ini agar lebih jelas, lengkap, dan punya contoh command yang valid
```

## Checklist kualitas

Sebelum publish, pastikan:

- command di docs cocok dengan source CLI
- tidak ada token/secret di contoh config
- halaman quickstart bisa diikuti user baru
- guide memiliki outcome yang jelas
- reference punya contoh command singkat
- semua link sidebar valid
- build VitePress sukses

## Verifikasi

Setiap perubahan docs harus diakhiri dengan:

```bash
cd docs-site
npm run docs:build
```

Untuk preview lokal:

```bash
cd docs-site
npm run docs:dev
```

## Kapan install Understand Anything?

Install Understand Anything eksternal berguna setelah docs dasar stabil, terutama untuk:

- audit gap dokumentasi lebih sistematis
- onboarding analysis dari sudut pandang tool eksternal
- visual knowledge graph tambahan
- validasi ulang struktur docs

Untuk membuat dokumentasi awal, Smara + source code + Graphify sudah cukup dan lebih aman.
