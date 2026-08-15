# Workspaces

Workspace membantu memisahkan konteks project, memory, dan sesi kerja. Ini penting jika kamu memakai Smara untuk beberapa repo atau client berbeda.

## Command

```bash
smara workspace list
smara workspace create <name>
smara workspace use <name>

smara explore [path]
```

## Kapan memakai workspace

Gunakan workspace saat:

- bekerja di beberapa repo,
- ingin memory per project tidak bercampur,
- menjalankan workflow berbeda per client,
- membuat skill/project context khusus,
- ingin audit atau docs generation untuk satu codebase tertentu.

## Membuat workspace

```bash
smara workspace create smara-cli
smara workspace use smara-cli
```

Setelah aktif, jalankan eksplorasi repo:

```bash
smara explore .
```

## Explore sebelum edit

`smara explore [path]` membantu agen memahami struktur project sebelum mengubah file.

Contoh prompt:

```text
Explore repo ini dulu secara read-only, lalu buat rencana update dokumentasi.
```

## Hubungan dengan memory

Workspace dapat membantu membatasi memory yang relevan. Contoh:

- workspace `smara-cli` menyimpan keputusan docs Smara,
- workspace `client-a` menyimpan konteks deployment client A,
- workspace `research` menyimpan hasil riset umum.

## Best practice

- Buat workspace per project besar.
- Simpan keputusan desain/arsitektur sebagai memory.
- Jangan mencampur secret antar workspace.
- Gunakan nama workspace yang jelas.
- Jalankan `explore` sebelum refactor atau docs generation besar.
