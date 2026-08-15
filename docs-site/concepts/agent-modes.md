# Agent Modes

Smara memakai mode agen untuk menyeimbangkan kecepatan, kontrol, dan keamanan.

## Ask

Mode `ask` cocok untuk tanya jawab cepat, penjelasan konsep, atau review ringan.

Karakteristik:

- Fokus pada jawaban teks.
- Minim tool call.
- Tidak cocok untuk perubahan file/server.

## Rush

Mode `rush` cocok untuk tugas kecil yang jelas dan aman.

Karakteristik:

- Langsung mengeksekusi tool yang diperlukan.
- Cocok untuk inspeksi cepat, build/test lokal, atau command non-destruktif.
- Perlu hati-hati untuk target server/production.

## Plan

Mode `plan` adalah mode paling aman untuk pekerjaan teknis.

Karakteristik:

- Eksplorasi read-only terlebih dahulu.
- Membuat rencana terstruktur.
- Menunggu approval user sebelum tool mutating/destructive/remote-write.
- Cocok untuk refactor, deployment, migrasi docs, dan operasi server.

## Rekomendasi

Gunakan aturan sederhana:

| Kebutuhan | Mode |
|---|---|
| Bertanya atau minta penjelasan | `ask` |
| Tugas kecil dan jelas | `rush` |
| Mengubah file, repo, server, deployment | `plan` |
