# Memory

Memory adalah basis pengetahuan Smara. Ia membantu agen mengingat preferensi, detail project, keputusan teknis, dan workflow berulang.

## Fitur utama

- Hybrid search: semantic + keyword.
- Kategori memori.
- Versioning dan rollback.
- Import/export backup.
- Memory graph: relasi antar memori.
- Cloud Memory: sync local-first antar device.

## Command dasar

```bash
smara memory search "cara deploy nextjs" --hybrid
smara memory save "Preferensi user: gunakan jawaban ringkas"
smara memory export backup.zip --format zip
smara memory import backup.zip
```

## Memory Graph

Hubungkan memori secara manual:

```bash
smara memory link 12 34 --relation refines --weight 0.8 --note "iterasi v2"
smara memory links 12
smara memory unlink 7
```

Bangun koneksi otomatis:

```bash
smara memory autolink --threshold 0.78 --top-k 5
```

Visualisasi:

```bash
smara memory graph
smara memory graph --port 7878
smara memory graph --export graph.json
```

## Best practice

- Simpan fakta jangka panjang, bukan noise percakapan sementara.
- Gunakan kategori untuk domain besar seperti `project`, `server`, `preference`, dan `release`.
- Backup sebelum migrasi atau cleanup besar.
