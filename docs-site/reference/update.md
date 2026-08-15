# Update Reference

`smara update` memperbarui binary Smara dari GitHub Releases.

```bash
smara update
smara update 1.4.0
smara update --version 1.4.0
smara update --no-restart
```

## Cara kerja

1. Mengambil release terbaru atau tag versi tertentu.
2. Mencari asset yang cocok dengan OS/arsitektur.
3. Mengunduh archive `.tar.gz` atau `.zip`.
4. Mengekstrak dan mengganti binary aktif.
5. Pada Linux, mencoba menjadwalkan restart systemd service yang memakai binary Smara.

## Flags

| Flag | Fungsi |
|---|---|
| `--version`, `-V` | Versi spesifik yang ingin diinstall. |
| `--no-restart` | Jangan restart otomatis service systemd. |

## Production checklist

- Backup binary lama sebelum update production.
- Jalankan `smara doctor` setelah update.
- Gunakan `--no-restart` jika restart service harus dijadwalkan manual.
- Cek release notes sebelum major update.
