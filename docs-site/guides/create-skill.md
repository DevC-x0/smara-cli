# Create a Skill

Skill adalah automation recipe yang bisa dijalankan ulang oleh Smara. Skill cocok untuk workflow berulang seperti deploy, backup, monitoring service, atau generate release notes.

## Kapan membuat skill

Buat skill jika task:

- terdiri dari beberapa langkah tool-call
- sering diulang
- punya parameter yang bisa diganti
- butuh urutan eksekusi konsisten

## Contoh permintaan

```text
buatkan skill untuk build docs-site lalu preview lokal
```

```text
simpan routine cek nginx + disk usage di VPS sebagai skill
```

## Format konsep

Skill biasanya berisi:

- nama skill
- deskripsi
- parameter opsional
- daftar step tool yang dijalankan berurutan
- tag kategori

## Best practice

- Gunakan nama kebab-case.
- Buat step kecil dan jelas.
- Tambahkan parameter untuk path, host, atau service name.
- Hindari hardcode secret.
- Verifikasi hasil di akhir workflow.
