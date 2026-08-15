# Deploy a Project

Smara bisa membantu deploy project secara bertahap: build lokal, upload asset, jalankan command di VPS, restart service, lalu cek log.

## Flow aman

1. Review struktur project.
2. Tentukan target server dan direktori deploy.
3. Build dan test lokal.
4. Backup versi remote jika perlu.
5. Upload artefak.
6. Restart service.
7. Cek health check dan log.
8. Siapkan rollback.

## Contoh prompt

```text
deploy docs-site ke VPS, tapi rencanakan dulu sebelum eksekusi
```

```text
build project ini, upload ke /var/www/app, lalu restart nginx setelah saya setuju
```

## Verification

- Build sukses.
- Service aktif.
- Endpoint dapat diakses.
- Log tidak menunjukkan error baru.

## Rollback

Untuk deploy penting, selalu simpan backup artefak atau release sebelumnya supaya bisa restore cepat jika health check gagal.
