# Multimodal Tools

Smara tidak hanya bekerja dengan teks. Beberapa command dan tool mendukung image, voice, desktop bridge, dan pointer automation untuk workflow yang lebih visual.

## Image generation

Command:

```bash
smara image "pastel green CLI docs hero" --out hero.png
```

Use case:

- membuat hero image docs,
- menghasilkan icon/illustration,
- mockup visual untuk landing page,
- asset cepat untuk prototype.

Tips:

- Simpan prompt dan output path agar reproducible.
- Jangan generate asset final brand tanpa review manual.
- Untuk docs, optimasi ukuran file sebelum publish.

## Voice

Command:

```bash
smara voice speak "Build docs selesai"
smara voice transcribe audio.wav
smara voice plan
```

Use case:

- transkripsi meeting atau voice note,
- membacakan ringkasan status,
- merancang voice interaction,
- accessibility workflow.

Safety:

- Jangan upload audio sensitif ke provider eksternal tanpa persetujuan.
- Simpan transkrip penting di lokasi aman.
- Review hasil transkripsi sebelum dipakai sebagai keputusan final.

## Desktop agent

Command:

```bash
smara desktop-agent
```

Desktop agent dipakai sebagai bridge untuk aktivitas lokal yang membutuhkan konteks desktop. Ketersediaan fitur tergantung OS dan permission.

Use case:

- menghubungkan workflow CLI dengan desktop,
- membantu inspeksi UI lokal,
- bridge untuk automation yang tidak murni terminal.

## Magic Pointer

Command:

```bash
smara magic-pointer
```

Magic Pointer ditujukan untuk interaksi pointer/desktop yang lebih visual. Gunakan hati-hati, terutama jika ada aplikasi production atau data sensitif terbuka.

## Best practice multimodal

- Mulai dari prompt yang spesifik.
- Simpan output di folder project yang jelas.
- Jangan overwrite asset lama tanpa backup.
- Untuk desktop/pointer, tutup aplikasi sensitif.
- Untuk voice/image provider eksternal, perhatikan data privacy.

## Workflow docs-site

Contoh prompt:

```text
Buat konsep hero image untuk Smara docs dengan warna lama: background #0a0a0a, accent #bef264, gaya terminal + knowledge graph.
```

Output asset bisa diletakkan di:

```text
docs-site/public/
```

Lalu referensikan dari Markdown atau theme VitePress.
