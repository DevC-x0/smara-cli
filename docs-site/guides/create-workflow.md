# Create a Workflow

Panduan ini menunjukkan cara membuat workflow Smara untuk pekerjaan multi-step. Workflow berguna saat tugas terlalu besar untuk satu skill, misalnya generate dokumentasi, audit server, atau release automation.

## 1. Tentukan outcome

Mulai dari hasil akhir yang jelas:

```text
Outcome: docs-site VitePress diperbarui dari source code aktual dan build sukses.
```

Hindari outcome yang terlalu samar seperti “rapikan project”.

## 2. Pecah menjadi fase

Contoh workflow dokumentasi:

```text
1. Scan source code read-only
2. Buat feature map
3. Buat docs gap report
4. Update Markdown
5. Update sidebar
6. Build VitePress
7. Ringkas perubahan
```

## 3. Buat workflow

```bash
smara workflow create docs-refresh
smara workflow show docs-refresh
```

Lalu isi blueprint sesuai format yang disediakan Smara pada versi aktif.

## 4. Jalankan workflow

```bash
smara workflow run docs-refresh
```

Untuk workflow berisiko, jalankan dari mode plan atau minta Smara menampilkan rencana dulu.

## 5. Tambahkan skill untuk step berulang

Jika ada step yang sering dipakai, jadikan skill.

Contoh:

```text
Buatkan skill untuk build docs-site: masuk folder docs-site, npm install jika perlu, npm run docs:build.
```

Workflow bisa memanggil skill tersebut sebagai bagian dari verification.

## 6. Verifikasi

Untuk docs-site:

```bash
cd docs-site
npm run docs:build
```

Untuk project Go:

```bash
gofmt -w ./...
go test ./...
```

Untuk server:

```bash
smara ssh exec prod "systemctl status app --no-pager"
```

## Template workflow sederhana

```text
Name: docs-refresh
Goal: keep Smara docs aligned with source code
Phases:
  - Analyze source
  - Compare docs
  - Update Markdown
  - Verify build
  - Summarize changes
Safety:
  - Do not edit non-doc bug fix files
  - Ask before remote write
Verification:
  - npm run docs:build in docs-site
```

## Safety checklist

Sebelum menjalankan workflow:

- Pastikan `git status --short` diketahui.
- Tandai file yang tidak boleh disentuh.
- Pisahkan read-only scan dan mutating edit.
- Jangan menjalankan command destructive tanpa approval.
- Selalu jalankan verification akhir.
