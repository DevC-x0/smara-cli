# Tools

Tools adalah kemampuan eksternal yang bisa dipakai Smara untuk membaca konteks, menjalankan command, mengedit file, mengakses server, mencari web, dan menganalisis codebase.

## Kenapa tools penting?

Tanpa tools, agen hanya bisa menjawab berdasarkan konteks percakapan. Dengan tools, Smara dapat:

- membaca file sebelum menyarankan perubahan
- menjalankan build/test
- mengedit dokumentasi atau source code
- mengelola VPS lewat SSH
- membaca dokumen dan screenshot
- membuat knowledge graph codebase
- menyimpan workflow berulang sebagai skill

## Kategori umum

| Kategori | Contoh kemampuan |
|---|---|
| File tools | baca, tulis, edit, cari file, grep. |
| Command tools | jalankan shell lokal, build, test, serve preview. |
| SSH tools | exec, view file, list dir, upload, download di VPS. |
| Web tools | search dan fetch halaman. |
| Document/media tools | baca PDF/DOCX/Markdown, OCR gambar, clipboard image. |
| Code intelligence | dependency analysis, call graph, LSP, Graphify. |
| Skill tools | membuat, menjalankan, dan mengelola automation skill. |
| MCP tools | menghubungkan tool eksternal dari MCP server. |

## Prinsip safety

Smara membedakan:

- **read-only tools** untuk eksplorasi aman
- **mutating tools** untuk aksi yang mengubah file/state lokal
- **remote-write tools** untuk aksi di server

Dalam Plan Mode, mutating dan remote-write dilakukan setelah approval user. Ini penting untuk mencegah agen langsung menghapus file, deploy, restart service, atau overwrite bug fix yang belum di-commit.

## Policy

Gunakan policy untuk mengatur izin tool automation:

```bash
smara policy list
smara policy set ssh_exec ask
smara policy set delete_file ask
smara policy check ssh_exec
```

Rekomendasi production:

- `ask` untuk SSH/deploy/delete
- `allow` untuk read-only search/list/view
- `deny` untuk tool yang tidak boleh dipakai di environment tertentu

## Tools untuk membuat dokumentasi

Workflow docs Smara biasanya memakai:

```text
read_file / grep_search / analyze_workspace
+ graphify_init / graphify_query
+ write_file / edit_file
+ run_command docs build
```

Contoh prompt:

```text
analisis dependency project ini dan jelaskan package pentingnya
```

```text
update docs-site/reference/cli-commands.md dari cmd/smara/*.go lalu jalankan build docs
```

```text
lihat log nginx di VPS, jangan restart dulu
```

## Lihat referensi tool

Untuk daftar kategori tool yang lebih lengkap, lihat [Tool List](/reference/tool-list).
