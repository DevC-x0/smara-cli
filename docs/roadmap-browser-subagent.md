# Roadmap Fitur Browser Subagent

## Context

Browser Subagent adalah fitur Smara/Antigravity untuk menjalankan browser automation dari prompt natural. Targetnya adalah membantu user melakukan E2E testing, visual checking, exploratory testing, screenshot capture, dan export laporan Markdown tanpa harus menulis script Playwright secara manual.

Contoh intent utama:

- Membuka `http://localhost:3000`, login sebagai admin, lalu screenshot dashboard.
- Membuka `http://localhost:5173`, mengecek navbar responsif di viewport mobile, lalu screenshot komponen navbar.
- Membuka halaman checkout, mencoba klik tombol `Bayar` tanpa mengisi form, mengecek error merah, lalu screenshot pesan error.

Outcome yang dituju:

- Smara mengenali kata kunci seperti `buka browser`, `gunakan browser subagent`, `ambil screenshot`, `testing E2E`, dan `periksa UI`.
- Browser Subagent bisa membuka URL lokal/remote, klik, input teks, submit form, menunggu perubahan UI, dan mengambil screenshot.
- Setiap run menghasilkan artifact berupa screenshot PNG, metadata JSON, dan laporan Markdown yang bisa diunduh atau dikirim lewat Smara Discord.

---

## Assumptions / Open Questions

### Assumptions

- Backend browser automation menggunakan Playwright/Chromium.
- Target awal adalah aplikasi lokal seperti React/Vue/Vite/Node server.
- Output minimal adalah screenshot `.png` dan laporan `report.md`.
- Untuk Discord/VPS, `localhost` mengarah ke mesin tempat bot berjalan, bukan laptop user.
- Browser Subagent berjalan dalam safe mode untuk mencegah aksi destruktif di production.

### Open Questions

- Browser Subagent akan aktif di CLI saja, Discord saja, atau keduanya?
- Apakah screenshot perlu dikirim otomatis sebagai attachment di Discord?
- Apakah perlu video recording selain screenshot?
- Apakah credential login perlu diambil dari env/secret seperti `ADMIN_USERNAME` dan `ADMIN_PASSWORD`?
- Apakah perlu allowlist domain untuk mencegah testing tidak sengaja pada production?

---

## Recommended Approach

Bangun fitur bertahap dalam lima milestone:

1. **MVP Screenshot** — buka URL, cek server, screenshot halaman, buat laporan Markdown.
2. **E2E Interaction Runner** — fill input, klik tombol, tunggu navigasi, screenshot hasil.
3. **Visual Checking** — viewport mobile/tablet/desktop, screenshot komponen, cek overflow.
4. **Exploratory Testing** — validasi form kosong, deteksi error, capture console/network issue.
5. **CLI + Discord Integration** — command `smara browser`, route intent di Discord, kirim artifact.

Pendekatan ini menjaga fitur tetap kecil, bisa diuji cepat, dan aman untuk digabung ke release bertahap.

---

## Milestone 1 — Browser Subagent MVP

### Goal

Smara bisa membuka browser ke URL yang diminta dan mengambil screenshot dasar.

### Fitur

- Deteksi intent browser dari prompt.
- Validasi URL lokal/remote.
- Cek apakah server reachable sebelum browser dijalankan.
- Launch Chromium headless/headful.
- Buka halaman target.
- Ambil screenshot full page.
- Simpan artifact ke folder run.
- Generate laporan Markdown.

### Artifact Structure

```txt
.smara/artifacts/browser-runs/<timestamp>/
├── screenshot-home.png
├── run.json
└── report.md
```

### Prompt Example

```txt
Buka browser ke http://localhost:3000 dan ambil screenshot.
```

### Acceptance Criteria

- Jika server tidak berjalan, Smara menampilkan error jelas:

```txt
Server http://localhost:3000 tidak bisa diakses. Jalankan development server dulu.
```

- Jika berhasil, screenshot tersimpan sebagai PNG.
- `report.md` berisi prompt asli, URL, waktu run, status, dan path screenshot.

---

## Milestone 2 — E2E Interaction Runner

### Goal

Browser Subagent bisa menjalankan skenario login dan interaksi UI dasar.

### Fitur

- Mengisi input berdasarkan label, placeholder, name, id, role, atau heuristic selector.
- Klik tombol/link berdasarkan teks.
- Tunggu navigasi selesai.
- Tunggu elemen tertentu muncul.
- Screenshot setelah step penting.
- Masking password di log dan report.

### Prompt Example

```txt
Gunakan browser subagent untuk membuka http://localhost:3000.
Tolong lakukan simulasi login dengan memasukkan username 'admin'
dan password 'password123'. Setelah berhasil masuk ke halaman dashboard,
ambil screenshot dan simpan hasilnya.
```

### Internal Task Plan Example

```json
{
  "url": "http://localhost:3000",
  "steps": [
    { "action": "goto", "target": "http://localhost:3000" },
    { "action": "fill", "target": "username", "value": "admin" },
    { "action": "fill", "target": "password", "value": "password123", "secret": true },
    { "action": "click", "target": "Login" },
    { "action": "waitFor", "target": "dashboard" },
    { "action": "screenshot", "name": "dashboard" }
  ]
}
```

### Acceptance Criteria

- Bisa menjalankan login pada aplikasi demo/local.
- Screenshot dashboard tersimpan.
- Report menyamarkan password, contoh: `password: ********`.
- Jika login gagal, report menyertakan screenshot state terakhir dan pesan error.

---

## Milestone 3 — Visual Checking Responsive UI

### Goal

Browser Subagent bisa membantu validasi tampilan responsif dan membuat screenshot komponen.

### Fitur

- Viewport preset:
  - Mobile: `375x812`
  - Tablet: `768x1024`
  - Desktop: `1440x900`
- Screenshot full page atau komponen tertentu.
- Deteksi elemen navbar:
  - `nav`
  - `[role="navigation"]`
  - class/id mengandung `navbar`, `nav`, atau `header`
- Deteksi horizontal overflow.
- Beri status `pass`, `needs_review`, atau `fail`.

### Prompt Example

```txt
Buka http://localhost:5173 di browser.
Tolong periksa apakah tata letak navbar sudah responsif di ukuran layar mobile.
Ambil screenshot pada komponen navbar tersebut agar saya bisa memvalidasi tampilannya.
```

### Artifact Structure

```txt
.smara/artifacts/browser-runs/<timestamp>/
├── navbar-mobile.png
├── navbar-tablet.png
├── navbar-desktop.png
├── visual-check.json
└── report.md
```

### Acceptance Criteria

- Screenshot navbar mobile berhasil dibuat jika navbar ditemukan.
- Jika navbar tidak ditemukan, Browser Subagent fallback ke screenshot full page.
- Report menyebutkan apakah ada horizontal overflow.

---

## Milestone 4 — Exploratory Testing & Bug Finding

### Goal

Browser Subagent bisa menjalankan eksplorasi bug sederhana berdasarkan instruksi natural.

### Fitur

- Navigasi ke halaman tertentu.
- Klik tombol tanpa mengisi form.
- Deteksi pesan error.
- Deteksi warna merah atau class error.
- Screenshot area error.
- Capture console errors.
- Capture failed network requests.

### Prompt Example

```txt
Tolong navigasikan browser ke halaman checkout di http://localhost:8000.
Cobalah klik tombol 'Bayar' tanpa mengisi form data diri.
Periksa apakah peringatan error merah muncul di layar,
lalu ambil screenshot dari pesan error tersebut.
```

### Internal Task Plan Example

```json
{
  "url": "http://localhost:8000/checkout",
  "steps": [
    { "action": "goto", "target": "http://localhost:8000/checkout" },
    { "action": "click", "target": "Bayar" },
    { "action": "assertVisible", "target": "error message" },
    { "action": "assertStyle", "target": "error message", "style": "red" },
    { "action": "screenshot", "name": "checkout-error" }
  ]
}
```

### Acceptance Criteria

- Tombol `Bayar` bisa diklik.
- Form kosong memunculkan validasi jika aplikasi benar.
- Screenshot error tersimpan.
- Report menyebutkan apakah error merah ditemukan.
- Jika tidak ditemukan, status `fail` dan report menyertakan screenshot bukti.

---

## Milestone 5 — Markdown Report Export

### Goal

Setiap run Browser Subagent menghasilkan laporan `.md` yang siap diunduh atau dikirim ke Discord.

### Fitur

- Generate `report.md` per run.
- Sertakan:
  - prompt asli
  - URL
  - waktu run
  - browser mode
  - viewport
  - langkah yang dijalankan
  - status tiap step
  - screenshot links
  - console errors
  - failed network requests
  - rekomendasi fix jika bug ditemukan
- Untuk Smara Discord, kirim screenshot `.png` dan `report.md` sebagai attachment.

### Report Example

```md
# Browser Subagent Report

## Prompt

Gunakan browser subagent untuk membuka http://localhost:3000...

## Environment

- URL: http://localhost:3000
- Browser: Chromium
- Mode: headless
- Viewport: 1440x900

## Steps

| Step | Action | Target | Status |
|---|---|---|---|
| 1 | goto | http://localhost:3000 | pass |
| 2 | fill | username | pass |
| 3 | fill | password | pass |
| 4 | click | Login | pass |
| 5 | screenshot | dashboard.png | pass |

## Screenshots

![Dashboard](./dashboard.png)

## Result

Status: pass
```

---

## Milestone 6 — CLI & Discord Integration

### Goal

Browser Subagent bisa dipakai dari Smara CLI dan Smara Discord.

### CLI Command Proposal

```bash
smara browser run "Buka http://localhost:3000 dan ambil screenshot"
```

```bash
smara browser run --url http://localhost:3000 --screenshot
```

```bash
smara browser e2e --spec browser-task.md
```

### Discord Behavior

Jika user di Discord menulis:

```txt
Gunakan browser subagent buka http://localhost:3000 dan ambil screenshot
```

Smara Discord akan:

1. Mengenali browser intent.
2. Menjalankan browser task.
3. Mengirim screenshot sebagai attachment.
4. Mengirim `report.md` sebagai attachment.

### Localhost Warning

Untuk Discord/VPS, tampilkan catatan:

```txt
Catatan: localhost mengarah ke mesin tempat bot berjalan, bukan perangkat user.
Jika target ada di laptop lokal, gunakan tunnel/public URL.
```

---

## Files / Tools Likely Needed

### Core Browser Package

```txt
internal/browser/
├── subagent.go
├── planner.go
├── runner.go
├── screenshots.go
├── report.go
├── server_check.go
└── types.go
```

### Playwright Driver

```txt
internal/browser/playwright/
├── driver.go
└── selectors.go
```

### Discord Integration

```txt
internal/platform/discord/
├── browser_intent.go
└── browser_artifacts.go
```

### CLI Integration

```txt
cmd/smara/
├── browser.go
└── root.go
```

### Documentation

```txt
docs/roadmap-browser-subagent.md
docs-site/public/roadmap-browser-subagent.md
docs-site/src/docs/UserGuide/BrowserSubagent.tsx
```

### Suggested Libraries / Tools

- Playwright
- Chromium
- Go test
- Markdown report writer
- PNG screenshot artifact manager
- Optional: axe-core for accessibility checks
- Optional: pixelmatch for visual regression

---

## Verification

### Automated Tests

1. **Intent detection test**
   - Prompt berisi `buka browser` harus masuk Browser Subagent.
   - Prompt berisi `ambil screenshot` harus masuk Browser Subagent.
   - Prompt biasa tidak boleh salah route.

2. **Planner test**
   - Prompt login menghasilkan steps:
     - goto
     - fill username
     - fill password
     - click login
     - screenshot dashboard

3. **Report test**
   - `report.md` berhasil dibuat.
   - Screenshot path masuk ke Markdown.
   - Password disamarkan.

4. **Server check test**
   - URL mati menghasilkan error jelas.
   - URL hidup lanjut ke browser runner.

5. **Discord artifact test**
   - Screenshot `.png` dikirim sebagai attachment.
   - Report `.md` dikirim sebagai attachment.

### Manual E2E Verification

#### Test 1 — Login E2E

```txt
Gunakan browser subagent untuk membuka http://localhost:3000.
Tolong lakukan simulasi login dengan memasukkan username 'admin'
dan password 'password123'. Setelah berhasil masuk ke halaman dashboard,
ambil screenshot dan simpan hasilnya.
```

Expected:

```txt
dashboard.png
report.md
status: pass / needs_review
```

#### Test 2 — Visual Navbar Mobile

```txt
Buka http://localhost:5173 di browser.
Tolong periksa apakah tata letak navbar sudah responsif di ukuran layar mobile.
Ambil screenshot pada komponen navbar tersebut agar saya bisa memvalidasi tampilannya.
```

Expected:

```txt
navbar-mobile.png
report.md
status: needs_review
```

#### Test 3 — Checkout Form Error

```txt
Tolong navigasikan browser ke halaman checkout di http://localhost:8000.
Cobalah klik tombol 'Bayar' tanpa mengisi form data diri.
Periksa apakah peringatan error merah muncul di layar,
lalu ambil screenshot dari pesan error tersebut.
```

Expected:

```txt
checkout-error.png
report.md
status: pass jika error muncul
status: fail jika error tidak muncul
```

---

## Risks / Rollback

### Risks

1. **Localhost ambiguity**
   - Di Discord/VPS, `localhost` bukan laptop user.
   - Mitigasi: tampilkan warning dan dukung tunnel/public URL.

2. **Selector tidak selalu akurat**
   - AI bisa salah memilih input/tombol.
   - Mitigasi: fallback ke role, label, placeholder, text, CSS heuristic, dan screenshot debug.

3. **Aksi destruktif**
   - Prompt bisa meminta klik bayar/delete/submit production.
   - Mitigasi: safe mode, domain allowlist, dan konfirmasi sebelum aksi sensitif.

4. **Credential exposure**
   - Password di prompt bisa bocor ke report.
   - Mitigasi: masking otomatis di logs/report.

5. **Browser dependency berat**
   - Playwright/Chromium butuh install besar.
   - Mitigasi: lazy install, command doctor, dan pesan error jelas.

6. **Flaky tests**
   - UI lambat atau animasi bisa membuat test gagal.
   - Mitigasi: wait strategy, retry terbatas, dan timeout configurable.

### Rollback

Jika fitur bermasalah:

- Matikan route Browser Subagent via config:

```txt
SMARA_BROWSER_SUBAGENT=false
```

- Pertahankan command manual screenshot jika masih aman.
- Hapus integrasi Discord intent tanpa menghapus core browser package.
- Rollback dependency Playwright jika build/deploy bermasalah.

---

## Priority Summary

| Priority | Milestone | Output |
|---|---|---|
| P0 | MVP Screenshot | URL check, screenshot, report.md |
| P1 | Login E2E | Fill, click, wait, dashboard screenshot |
| P1 | Visual Checking | Mobile viewport, navbar screenshot, overflow check |
| P2 | Exploratory Testing | Form validation, error screenshot, console/network capture |
| P2 | CLI Integration | `smara browser run` command |
| P2 | Discord Integration | Screenshot + Markdown attachment |
| P3 | Advanced Visual Regression | Baseline comparison, diff image, threshold |
| P3 | Accessibility Check | axe-core report and suggestions |

---

## Status

Roadmap ini siap dipakai sebagai acuan implementasi Browser Subagent untuk Smara CLI dan Smara Discord.
