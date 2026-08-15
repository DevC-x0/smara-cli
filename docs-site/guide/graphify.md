# Graphify Knowledge Graph (safishamsi/graphify)

Graphify mengubah codebase menjadi knowledge graph yang bisa ditelusuri, di-query, dan diekspor. Project ini menggunakan [safishamsi/graphify](https://github.com/safishamsi/graphify) — Python-based tool dengan dukungan 23+ bahasa pemrograman.

## Fitur Utama

- **Multi-language parsing** (Go, JavaScript, TypeScript, Python, Rust, Java, C++, dll) menggunakan tree-sitter
- **Natural language query** dengan BFS/DFS traversal
- **Community detection** (Leiden algorithm)
- **Watch mode** untuk auto-rebuild saat file berubah
- **Merge multiple repos** ke satu cross-repo graph
- **No LLM needed** untuk update (AST-only)

## Setup

Graphify sudah terinstall via uv:
```bash
which graphify  # /home/cahya/.local/bin/graphify
```

## Build/Update Graph

**Initial build** (pakai LLM untuk semantic extraction):
```bash
graphify .
```

**Update setelah ubah kode** (no LLM, fast):
```bash
graphify update .
```

**Watch mode** (auto-rebuild saat file berubah):
```bash
graphify watch .
```

## Query

**Natural language query:**
```bash
graphify query "how does auth work"
graphify query "auth flow" --budget 1000
graphify query "database connections" --dfs  # depth-first traversal
```

**Pathfinding:**
```bash
graphify path "Handler" "Database"
```

**Explain node:**
```bash
graphify explain "LoadCustomWorkflow"
```

## Output Files

Setelah build/update, graphify menghasilkan:

- `graphify-out/graph.json` — data graph utama (nodes, edges, metadata)
- `graphify-out/GRAPH_REPORT.md` — laporan struktur, god nodes, communities
- `graphify-out/graph.html` — interactive visualization (jika <5000 nodes)
- `graphify-out/.graphify_*.json` — cache AST extraction

## Advanced Features

**Add external content:**
```bash
graphify add "https://example.com/paper.pdf" --author "John Doe"
```

**Merge multiple repos:**
```bash
graphify merge-graphs repo1/graphify-out/graph.json repo2/graphify-out/graph.json --out merged.json
```

**Tree visualization:**
```bash
graphify tree --output tree.html
```

**Git merge driver** (untuk graph.json conflicts):
```bash
graphify merge-driver <base> <current> <other>
```

## Auto-Update dengan Git Hook

Pasang post-commit hook untuk auto-update graph setiap commit:

```bash
cp scripts/git-hook-post-commit-graphify.sh .git/hooks/post-commit
chmod +x .git/hooks/post-commit
```

Script akan jalankan `graphify update .` secara otomatis (silent, no LLM).

## Agent Integration

Graphify terintegrasi dengan berbagai AI coding assistants:
- Claude Code
- Codex
- OpenCode
- Cursor
- Gemini CLI
- Aider
- Kiro
- Dan lainnya

Install skill:
```bash
graphify install --platform claude
graphify install --platform codex
graphify install --platform cursor
```

## Current Stats (2026-06-20)

- **6220 nodes** (functions, types, variables, concepts)
- **11270 edges** (calls, imports, contains, semantically_similar_to)
- **361 communities** detected
- **676 files** parsed
- Built with tree-sitter (multi-language)

## Tips

- Gunakan `--budget N` untuk limit output token saat query
- Gunakan `--dfs` untuk traversal lebih dalam (default BFS)
- Untuk graph >5000 nodes, HTML viz tidak di-generate (gunakan `--no-viz` atau set `GRAPHIFY_VIZ_NODE_LIMIT`)
- Set `MOONSHOT_API_KEY` untuk pakai Kimi K2.6 (3x lebih murah dari GPT-4)

## Troubleshooting

**Graph tidak ter-update:**
```bash
graphify update . --force  # overwrite meskipun nodes lebih sedikit
```

**Check freshness:**
```bash
graphify check-update .
```

**Rebuild clustering only:**
```bash
graphify cluster-only . --no-viz
```

## Resources

- GitHub: https://github.com/safishamsi/graphify
- Documentation: https://github.com/safishamsi/graphify#readme
