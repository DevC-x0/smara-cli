# Docs Generation Example

Contoh ini menunjukkan workflow membuat dokumentasi Smara dengan gabungan analisis codebase, knowledge graph, dan VitePress.

## Goal

Membuat halaman baru untuk fitur Smara berdasarkan source code dan README.

## Prompt

```text
analisis fitur skill di repo ini, lalu buat halaman VitePress di docs-site/core-concepts/skills.md dengan contoh penggunaan dan best practice
```

## Expected workflow

1. Smara membaca file terkait skill.
2. Smara mencari README atau docs lama.
3. Smara menyusun outline.
4. Smara menulis Markdown.
5. Smara menjalankan build docs.

## Verification

```bash
cd docs-site
npm run docs:build
```

## Follow-up prompt

```text
cek apakah halaman skill sudah mencakup create, run, dependency, parameter, dan keamanan secret
```

## Pattern reusable

Pola yang sama bisa dipakai untuk:

- memory docs
- MCP docs
- SSH guide
- CLI command reference
- troubleshooting guide
