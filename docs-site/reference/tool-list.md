# Tool List

Halaman ini merangkum tool bawaan Smara yang umum tersedia untuk agen. Daftar aktual dapat berubah sesuai mode, policy, environment lokal, dan MCP server yang terkoneksi.

## Prinsip tool Smara

Smara membagi tool berdasarkan risiko:

- **Read-only**: aman untuk eksplorasi, misalnya baca file, list folder, search, analisis dependency.
- **Local mutating**: mengubah workspace lokal, misalnya tulis/edit/hapus file atau menjalankan command build yang menghasilkan artefak.
- **Remote mutating**: mengubah state server, misalnya upload file, restart service, deploy, atau command SSH.

Dalam **Plan Mode**, Smara menyusun rencana dulu dan meminta approval sebelum tool mutating/remote-write.

## File dan workspace

| Tool | Kegunaan |
|---|---|
| `view_file` / `read_file` | Membaca isi file sebelum edit. |
| `write_file` | Membuat atau menimpa file. |
| `edit_file` | Mengganti bagian spesifik file. |
| `delete_file` | Menghapus file. |
| `list_dir` | Melihat isi folder. |
| `search_path` | Mencari file/folder berdasarkan nama. |
| `grep_search` | Mencari teks di repo. |
| `analyze_workspace` | Ringkasan struktur workspace. |

## Command lokal dan preview

| Tool | Kegunaan |
|---|---|
| `run_command` | Menjalankan shell command lokal. |
| `serve_project` | Menjalankan server preview lokal untuk project web/static. |

Contoh penggunaan: `npm run build`, `go test ./...`, atau preview `docs-site` VitePress.

## SSH / VPS

| Tool | Kegunaan |
|---|---|
| `ssh_manage` | Tambah/hapus/list host SSH tersimpan. |
| `ssh_exec` | Jalankan command di server remote. |
| `ssh_view_file` | Baca file remote. |
| `ssh_list_dir` | List folder remote. |
| `ssh_upload` | Upload file lokal ke server. |
| `ssh_download` | Download file dari server. |

Gunakan dengan hati-hati. Untuk production, prefer mode `plan` dan policy `ask`.

## Web, dokumen, dan media

| Tool | Kegunaan |
|---|---|
| `web_search` | Pencarian web. |
| `web_fetch` | Ambil dan bersihkan konten halaman web. |
| `read_document` | Ekstrak teks PDF/DOCX/ODT/RTF/Markdown. |
| `analyze_image` | Metadata dan OCR gambar. |
| `clip_paste_image` | Ambil gambar dari clipboard. |
| `clip_copy_image` | Salin gambar ke clipboard. |
| `generate_image` | Generate gambar dari prompt. |
| `export_data` | Export data ke CSV/JSON/Markdown/PDF. |

## Code intelligence dan graph

| Tool | Kegunaan |
|---|---|
| `lsp_hover` | Dokumentasi/type symbol pada posisi tertentu. |
| `lsp_definition` | Lompat ke definisi symbol. |
| `lsp_references` | Cari semua reference symbol. |
| `lsp_document_symbols` | Daftar symbol dalam file. |
| `analyze_dependencies` | Map import dan dependency internal/eksternal. |
| `generate_call_graph` | Outline call graph statis. |
| `graphify_init` | Bangun knowledge graph dari codebase Go. |
| `graphify_query` | Query knowledge graph. |

Tool ini mendukung workflow **Understand Anything style**: memahami struktur repo sebelum membuat docs atau refactor.

## Binary dan security analysis

| Tool | Kegunaan |
|---|---|
| `analyze_binary` | Deteksi format/arsitektur/entropy binary secara read-only. |
| `extract_strings` | Ekstrak string dari binary/dokumen. |
| `scan_signature` | YARA-lite matching untuk pattern byte/string/regex. |

## Skills, memory, dan automation

| Tool | Kegunaan |
|---|---|
| `skill_list` | Daftar skill tersimpan. |
| `skill_run` | Jalankan skill. |
| `skill_create` | Buat skill automation baru. |
| `skill_delete` | Hapus skill jika user eksplisit meminta. |
| `remember` | Simpan memori jangka panjang. |
| `search_memories` | Cari memori tersimpan. |
| `schedule_reminder` | Jadwalkan reminder/nudge. |
| `planning_template` | Scaffold planning terstruktur. |

## MCP

| Tool | Kegunaan |
|---|---|
| `connect_mcp` | Hubungkan MCP server local/remote. |
| `disconnect_mcp` | Putuskan koneksi MCP server. |

MCP memperluas tool Smara dengan kemampuan dari server eksternal.

## Iteration budget dan user profile

| Tool | Kegunaan |
|---|---|
| `iteration_budget_status` | Cek kuota iterasi tool turn saat ini. |
| `request_iteration_budget` | Minta tambahan iterasi untuk task besar. |
| `user_model` | Baca/update preferensi user. |

## Rekomendasi untuk dokumentasi

Untuk membuat docs dari repo Smara tanpa install Understand Anything eksternal, pakai kombinasi:

```text
read_file / grep_search / analyze_dependencies
+ graphify_init / graphify_query
+ write_file / edit_file
+ npm run docs:build
```

Output final tetap Markdown di `docs-site/**/*.md` dan dipublish dengan VitePress.
