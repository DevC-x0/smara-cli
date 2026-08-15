# React Deploy Example

Contoh workflow deploy aplikasi React dengan Smara.

## Prompt

```text
rencanakan deploy React app ini ke VPS: build lokal, upload dist, restart nginx, lalu cek endpoint
```

## Steps umum

1. Install dependency.
2. Build app.
3. Cek output `dist/`.
4. Backup folder remote.
5. Upload build artifact.
6. Reload nginx.
7. Cek status dan endpoint.

## Verification

```bash
npm run build
```

Lalu cek URL aplikasi dari browser atau `curl`.

## Rollback

Simpan backup release sebelumnya di server agar bisa restore jika build baru bermasalah.
