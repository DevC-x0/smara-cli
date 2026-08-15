# Graphify Reference

Graphify adalah subsystem Smara untuk membuat **knowledge graph dari codebase Go**. Ia membantu agen memahami relasi package, function, type, method, file, dan dependency internal sebelum membuat rencana atau dokumentasi.

## Kapan dipakai

Gunakan Graphify saat ingin:

- memahami repo Go yang besar;
- mencari fungsi/type terkait fitur tertentu;
- melihat jalur relasi antar node source code;
- membuat onboarding map untuk developer baru;
- mencari gap dokumentasi berdasarkan struktur aktual codebase;
- mengekspor graph ke format visual atau graph database.

## Command overview

```bash
smara graphify init [path] [--name smara] [--path .]
smara graphify query "memory skill mcp" --name smara --depth 2 [--budget 2000]
smara graphify explain "NodeName" --name smara --depth 2
smara graphify path "FromNode" "ToNode" --name smara
smara graphify export --name smara --format json --out graph.json
smara graphify list
smara graphify delete smara
```

## Init graph

```bash
smara graphify init . --name smara
```

Yang terjadi:

1. Smara mem-parse codebase Go pada path target.
2. Node dibuat dari simbol penting seperti package, file, type, function, method.
3. Edge dibuat dari relasi seperti contains, calls, imports, implements, references.
4. Community/cluster dihitung agar node terkait bisa dikelompokkan.
5. Graph disimpan ke SQLite memory database Smara.

::: tip
Untuk repo besar, mulai dari root module Go. Hindari menjalankan dari folder yang berisi vendor/build output besar jika tidak diperlukan.
:::

## Query graph

```bash
smara graphify query "provider login model config" --name smara --depth 2
```

Flag penting:

| Flag | Fungsi |
|---|---|
| `--name` | Nama graph yang sudah disimpan. Required untuk query. |
| `--depth` | Kedalaman neighborhood node hasil query. Default `2`. |
| `--budget` | Batasi output compact agar cocok dimasukkan ke prompt LLM. |

Contoh untuk dokumentasi:

```bash
smara graphify query "memory graph cloud sync" --name smara --depth 2 --budget 2500
smara graphify query "skill create run registry" --name smara --depth 2 --budget 2500
smara graphify query "ssh upload download exec" --name smara --depth 2 --budget 2000
```

## Explain node

```bash
smara graphify explain "runProviderSet" --name smara --depth 1
```

Gunakan untuk melihat konteks sekitar satu node: file asal, caller/callee, type terkait, dan relasi sekitar.

## Shortest path

```bash
smara graphify path "memory.Store" "cloud sync" --name smara
```

Berguna untuk menemukan hubungan konseptual antar fitur, misalnya:

- provider config → LLM runtime;
- memory store → cloud sync;
- skill format → skill runner;
- CLI command → internal implementation.

## Export

```bash
smara graphify export --name smara --format json --out graphify-out/smara.json
smara graphify export --name smara --format svg --out graphify-out/smara.svg
smara graphify export --name smara --format graphml --out graphify-out/smara.graphml
smara graphify export --name smara --format neo4j --out graphify-out/smara.cypher
```

Format:

| Format | Use case |
|---|---|
| `json` | Dipakai ulang oleh docs generator, Vue component, atau script analisis. |
| `svg` | Visual statis untuk README/docs. |
| `graphml` | Import ke yEd, Gephi, atau tool graph umum. |
| `neo4j` | Eksperimen query graph database. |

## Workflow untuk Smara docs

```text
1. smara graphify init . --name smara
2. Query area fitur: memory, skills, MCP, SSH, provider, web, workflow.
3. Ringkas node penting menjadi outline docs.
4. Cocokkan outline dengan sidebar VitePress.
5. Update Markdown.
6. Build docs-site.
```

Contoh prompt ke Smara:

```text
Gunakan graphify untuk mencari node terkait Cloud Memory, lalu update docs-site/guides/cloud-memory.md berdasarkan hasilnya. Jangan ubah file non-docs.
```

## Batasan

- Graphify saat ini paling kuat untuk codebase Go.
- Hasil graph adalah bantuan pemahaman, bukan pengganti review source code.
- Nama node bisa ambigu; gunakan `explain` atau `path` untuk validasi.
- Untuk source TypeScript/React, gunakan kombinasi `grep_search`, `analyze_dependencies`, dan dokumentasi manual.

## Best practice

- Pakai nama graph eksplisit seperti `smara` atau `smara-backend`.
- Export JSON saat akan dipakai oleh docs atau visualizer.
- Gunakan `--budget` sebelum memasukkan hasil ke LLM.
- Jalankan ulang `init` setelah refactor besar.
- Simpan insight penting ke memory atau docs agar tidak hanya tersimpan di graph lokal.
