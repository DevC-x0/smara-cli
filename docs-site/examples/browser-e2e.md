# Browser E2E Example

Contoh ini menunjukkan cara memakai browser automation Smara untuk mengecek docs-site atau aplikasi web.

## Goal

Memastikan homepage, sidebar, dan beberapa halaman utama bisa dibuka setelah build/deploy.

## 1. Jalankan app lokal

Untuk VitePress docs-site:

```bash
cd docs-site
npm run docs:dev
```

Catat URL yang muncul, misalnya `http://localhost:5173`.

## 2. Buat spec

`browser-task.md`:

```md
# Docs Site Smoke Test

Target: http://localhost:5173

Steps:
1. Open homepage.
2. Verify the Smara CLI title is visible.
3. Open Getting Started > Quickstart.
4. Open Guides > Knowledge Graph.
5. Open Reference > CLI Commands.
6. Report broken navigation, visual glitches, or console errors.

Constraints:
- Read-only only.
- Do not edit content from browser.
```

## 3. Jalankan E2E

```bash
smara browser e2e --spec browser-task.md
```

Atau prompt langsung:

```bash
smara browser run "cek docs-site lokal di http://localhost:5173, buka homepage, quickstart, knowledge graph, dan CLI commands"
```

## 4. Interpretasi hasil

Jika ada masalah:

- broken link: update `.vitepress/config.ts` atau link Markdown,
- layout gelap tidak readable: update `theme/custom.css`,
- console error: cek component/custom HTML,
- halaman 404: cek nama file dan clean URL.

## 5. Verification akhir

Tetap jalankan build static:

```bash
cd docs-site
npm run docs:build
```

Browser E2E dan build saling melengkapi: build memastikan static generation sukses, browser E2E memastikan flow pengguna bisa dipakai.
