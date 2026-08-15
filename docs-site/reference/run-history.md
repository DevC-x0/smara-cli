# Run History Reference

Run history membantu melihat sesi/eksekusi Smara sebelumnya, melakukan debugging, dan replay workflow tertentu.

## Command

```bash
smara run history
smara run show <id>
smara run replay <id>
smara run timeline <id>
```

## Use case

Gunakan run history untuk:

- melihat apa yang dilakukan agen,
- audit tool call dan keputusan,
- mencari error pada workflow panjang,
- replay proses yang berhasil,
- membuat dokumentasi dari sesi nyata.

## Melihat daftar run

```bash
smara run history
```

Cari run berdasarkan waktu, project, atau deskripsi.

## Melihat detail run

```bash
smara run show <id>
```

Detail biasanya mencakup prompt, step penting, tool call, output, dan status akhir sesuai data yang tersedia.

## Replay run

```bash
smara run replay <id>
```

Gunakan replay dengan hati-hati jika run lama berisi operasi mutating seperti edit file, upload, deploy, atau restart service.

Rekomendasi:

- replay read-only aman,
- replay mutating sebaiknya di mode plan,
- cek `git status` sebelum replay yang mengubah workspace,
- jangan replay deploy production tanpa maintenance window.

## Timeline

```bash
smara run timeline <id>
```

Timeline membantu memahami urutan eksekusi. Ini berguna untuk postmortem bug atau membuat runbook.

## Docs workflow

Run history dapat dipakai untuk membuat docs dari aktivitas nyata:

1. Jalankan workflow manual.
2. Cek `smara run timeline <id>`.
3. Ringkas step yang berhasil.
4. Tulis sebagai guide atau example.
5. Verifikasi command di docs.
