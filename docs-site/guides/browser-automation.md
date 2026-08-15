# Browser Automation

Smara memiliki browser automation untuk eksplorasi web, testing end-to-end ringan, dan validasi UI. Fitur ini berguna ketika agen perlu melihat halaman, mengikuti flow, atau membuat laporan browser.

## Command

```bash
smara browser run "cek halaman login dan laporkan error visual"
smara browser e2e --spec browser-task.md
```

Command aktual tersedia di:

```text
cmd/smara/browser.go
internal/browser/
```

## Kapan dipakai

Gunakan browser automation untuk:

- smoke test halaman web,
- cek link dan navigasi,
- validasi form sederhana,
- reproduksi bug UI,
- mengambil screenshot/report untuk investigasi,
- membandingkan hasil deploy.

## Contoh prompt

```text
Buka docs-site lokal, cek homepage, klik Guides, lalu laporkan link yang broken atau layout yang terlihat aneh.
```

```text
Jalankan browser E2E untuk flow login, tapi jangan submit data production.
```

## Spec file E2E

Contoh `browser-task.md`:

```md
# Browser Task

Target: http://localhost:5173

Steps:
1. Open homepage.
2. Verify hero text is visible.
3. Click Guides.
4. Open Knowledge Graph page.
5. Report visual/layout issues.

Constraints:
- Do not enter real credentials.
- Do not modify production data.
```

Lalu jalankan:

```bash
smara browser e2e --spec browser-task.md
```

## Output yang diharapkan

Tergantung versi dan environment, browser automation dapat menghasilkan:

- ringkasan hasil,
- screenshot,
- diagnostics,
- daftar error console/network,
- rekomendasi fix.

## Safety

- Jangan gunakan credential asli kecuali benar-benar perlu dan environment aman.
- Hindari destructive action di production.
- Untuk web admin, gunakan akun test.
- Untuk checkout/payment, gunakan sandbox.
- Jika browser automation perlu write action, minta plan dulu.

## Workflow dengan docs-site

```bash
cd docs-site
npm run docs:dev
```

Lalu minta Smara:

```text
Gunakan browser automation untuk cek docs-site lokal: homepage, sidebar, search, dan halaman Knowledge Graph.
```

Ini melengkapi verification `npm run docs:build` karena build sukses belum tentu layout atau navigasi terasa bagus.
