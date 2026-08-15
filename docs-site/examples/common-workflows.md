# Common Workflows

Halaman ini merangkum workflow yang paling sering dipakai dengan Smara. Gunakan sebagai prompt template atau dasar membuat skill/workflow.

## 1. Analisis repo sebelum refactor

```text
Tolong analisis struktur repo ini secara read-only dan buat rencana refactor command package.
```

Gunakan mode `plan`. Smara akan eksplorasi read-only, menyusun rencana, lalu meminta approval sebelum edit.

## 2. Buat skill untuk workflow berulang

```text
Buatkan skill untuk build, test, dan package release Smara.
```

Skill cocok jika langkahnya sering diulang dan punya parameter.

## 3. Buat workflow multi-step

```text
Buat workflow docs-refresh: scan source, buat gap report, update Markdown, update sidebar, build docs, lalu ringkas perubahan.
```

Workflow cocok jika task punya banyak fase, role, atau perlu resume.

## 4. Monitor VPS

```text
Cek status docker, disk usage, dan service nginx di vps prod.
```

Smara dapat memakai SSH tools untuk inspect server. Untuk perubahan config, gunakan mode `plan`.

## 5. Generate knowledge graph

```bash
smara graphify init . --name smara
smara graphify query "agent execution flow" --name smara --depth 3
```

Pakai hasil graph sebagai bahan dokumentasi, onboarding, atau refactor.

## 6. Browser smoke test

```bash
smara browser run "cek docs-site lokal di http://localhost:5173: homepage, quickstart, knowledge graph, dan CLI commands"
```

Atau buat spec:

```bash
smara browser e2e --spec browser-task.md
```

## 7. Backup dan restore memory

```bash
smara memory export smara-memory.zip --format zip
smara memory import smara-memory.zip
```

## 8. Jadwalkan automation

```bash
smara schedule add "daily at 09:00" server-health-check
smara schedule list
smara schedule daemon
```

Pastikan workflow yang dijadwalkan sudah dites manual.

## 9. Jalankan bot platform

```bash
smara serve --platform telegram --mode plan
```

Perintah bot umum:

- `/ask <prompt>`
- `/mode <ask|rush|plan>`
- `/mcp`
- `/clear`

## 10. Generate docs dari source code

```text
Scan source Smara, bandingkan dengan docs-site, buat gap report, lalu update halaman Markdown yang prioritas. Jangan sentuh file bug fix non-docs.
```

Verification:

```bash
cd docs-site
npm run docs:build
```


## 11. Monitoring VPS sebagai skill

Lihat contoh lengkap: [VPS Monitoring Skill](/examples/vps-monitoring-skill).

## 12. Release checklist

Lihat checklist lengkap: [Release Checklist](/examples/release-checklist).

## 13. Cloud Memory sync

Lihat contoh aman: [Cloud Memory Sync](/examples/cloud-memory-sync).
