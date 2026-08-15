# Agent Mode

Agent Mode adalah cara Smara bekerja sebagai agen developer: memahami tujuan, memakai tools, menjalankan command, mengedit file, melakukan verifikasi, dan melaporkan hasil.

## Kapan dipakai

Gunakan Agent Mode untuk tugas seperti:

- memperbaiki bug
- menambah fitur
- membuat dokumentasi
- menjalankan build/test
- deploy ke VPS
- membuat skill otomasi

## Plan-first safety

Untuk aksi yang berisiko atau mengubah state, Smara menyusun rencana terlebih dahulu:

1. memahami konteks
2. eksplorasi read-only bila perlu
3. menjelaskan langkah
4. meminta approval
5. eksekusi bertahap setelah disetujui
6. verifikasi hasil

## Contoh

```text
buatkan halaman docs baru untuk fitur memory, lalu build test
```

```text
cek service di VPS dan perbaiki kalau crash
```
