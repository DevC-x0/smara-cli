# Knowledge Graph

Knowledge graph membantu memahami Smara sebagai sistem yang saling terhubung: agent mode, memory, skills, tools, MCP, SSH, web UI, dan docs generation.

<div class="smara-graph glow-border">
  <span class="smara-node primary">Smara</span>
  <span class="smara-node">Agent</span>
  <span class="smara-node">Memory</span>
  <span class="smara-node">Skills</span>
  <span class="smara-node">Tools</span>
  <span class="smara-node">MCP</span>
  <span class="smara-node">SSH</span>
  <span class="smara-node">Docs</span>
</div>

## Konsep

Dalam konteks Smara docs:

```text
Understand Anything = analisis struktur + knowledge graph
VitePress = doc site utama
Smara/LLM = generate dan refine konten Markdown
```

Graph bukan hanya visual. Graph dipakai untuk menjawab pertanyaan seperti:

- fitur apa yang belum punya dokumentasi?
- command mana yang terkait dengan memory?
- file mana yang mengimplementasikan SSH remote control?
- workflow apa yang perlu contoh end-to-end?
- konsep mana yang harus muncul di onboarding user baru?

## Graph di Smara

Smara punya beberapa bentuk graph:

| Graph | Fungsi |
|---|---|
| **Graphify** | Knowledge graph dari codebase Go. Cocok untuk memahami package, function, type, dan relasi source code. |
| **Memory Graph** | Relasi antar memory user seperti `refines`, `supports`, dan `follows`. |
| **Skill Tree** | Hierarki skill, dependency edge, analytics, dan refinement workflow. |
| **Docs Graph** | Peta konseptual docs: fitur → guide → reference → example. |

## Menggunakan Graphify

Smara memiliki tooling Graphify untuk membuat knowledge graph dari codebase Go.

```bash
smara graphify init . --name smara
smara graphify query "skills memory mcp ssh" --name smara --depth 2
smara graphify explain "Skill" --name smara --depth 1
smara graphify export --name smara --format json
```

Jika sudah ada export graph:

```text
graphify-out/graph.json
graphify-out/graph.html
smara.graphml
```

File tersebut dapat menjadi bahan halaman docs, onboarding map, atau visualisasi interaktif.

## Workflow docs berbasis graph

```text
1. Scan source code
2. Build graph
3. Query fitur utama
4. Cocokkan dengan docs-site
5. Temukan gap docs
6. Tulis/refine Markdown
7. Build VitePress
```

Contoh prompt:

```text
pakai graphify untuk mencari relasi antara memory, skill, dan mcp, lalu update docs-site/guides/knowledge-graph.md
```

```text
bandingkan fitur di README dengan docs-site sidebar, buat daftar gap dokumentasi
```

## Roadmap interaktif

Tahap awal docs memakai visual statis agar cepat dan ringan. Tahap berikutnya bisa ditingkatkan menjadi:

- Vue component interaktif di VitePress
- import `graphify-out/graph.json`
- filter node berdasarkan kategori fitur
- klik node untuk membuka halaman docs terkait
- onboarding map ala Understand Anything demo

## Rekomendasi implementasi interaktif

Untuk menjaga VitePress tetap ringan:

1. Mulai dari component Vue kecil tanpa dependency berat.
2. Load graph JSON secara lazy hanya di halaman graph.
3. Batasi node awal ke fitur utama, bukan seluruh AST.
4. Tambahkan search/filter kategori.
5. Simpan link ke halaman docs di metadata node.
