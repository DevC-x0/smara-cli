# Analytics Reference

`smara analytics` menampilkan ringkasan penggunaan token, cost, model, request, prompt, dan skill.

```bash
smara analytics
smara analytics --days 7
smara analytics --days 30
```

## Data source

Smara membaca analytics dari path metrics lokal. Jika `db_path` kosong, fallback legacy memakai:

```text
./usage_analytics.jsonl
```

Pada konfigurasi normal, analytics mengikuti lokasi database/config Smara.

## Output

Ringkasan analytics dapat mencakup:

- total prompt/request,
- token input/output,
- estimasi cost,
- model yang dipakai,
- grafik/summary token harian,
- skill yang paling sering dipakai.

## Kapan dipakai

Gunakan analytics untuk:

- mengecek biaya model cloud,
- melihat workflow yang paling banyak berjalan,
- memilih provider/model yang lebih efisien,
- membuat release/growth report,
- audit penggunaan agent dalam tim.

## Best practice

- Gunakan `--days 7` untuk review mingguan.
- Gunakan `--days 30` untuk report bulanan.
- Jangan commit file analytics mentah jika mengandung prompt sensitif.
- Untuk tim, ekspor ringkasan saja, bukan raw prompt log.
