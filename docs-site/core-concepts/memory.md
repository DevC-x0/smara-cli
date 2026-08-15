# Memory

Memory adalah basis pengetahuan Smara. Ia menyimpan konteks jangka panjang agar agen tidak perlu memulai dari nol di setiap sesi: preferensi user, detail project, keputusan teknis, catatan server, release notes, dan workflow yang sering diulang.

Smara memakai pendekatan **local-first**. Database utama berada di mesin pengguna, lalu dapat ditambah Cloud Memory jika ingin sinkron antar device.

## Peran memory dalam workflow agen

Memory dipakai untuk:

- mengingat preferensi komunikasi dan mode kerja;
- menyimpan fakta project seperti host VPS, path repo, command build, dan keputusan arsitektur;
- memberi konteks tambahan sebelum agen membuat rencana atau menjalankan tool;
- menghubungkan pengalaman lama dengan task baru melalui search dan memory graph;
- mendukung dokumentasi, misalnya menyimpan catatan “fitur ini sudah didokumentasikan di halaman X”.

## Storage

Secara default Smara menyimpan memory di SQLite:

```yaml
db_path: ~/.smara/memory.db
```

Saat Cloud Memory aktif, command memory tetap memakai interface yang sama, tetapi store dibuka dengan sinkronisasi cloud-aware.

## Command dasar

```bash
smara memory save "Preferensi user: gunakan jawaban ringkas" --tags preference,user
smara memory list --limit 20
smara memory search "cara deploy docs" --hybrid
smara memory export backup.zip --format zip
smara memory import backup.zip
smara memory clear --older-than 90d --dry-run
```
```
Beberapa command mendukung filter seperti tags, source, category, tanggal, dan sort order.

## Kategori dan retensi

Gunakan kategori untuk memisahkan domain besar:

- `project` — detail repo, command build, arsitektur;
- `server` — host, service, path deploy;
- `preference` — gaya jawaban, bahasa, risk tolerance;
- `release` — catatan versi dan changelog;
- `docs` — keputusan dokumentasi.

Untuk memory yang hanya berlaku sementara, gunakan expiry/retention jika tersedia agar database tetap bersih.

## Memory Graph

Memory graph menghubungkan satu memori dengan memori lain. Ini berguna untuk melacak keputusan, refinement, dependensi, atau catatan yang saling terkait.

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

## Cloud Memory

Cloud Memory menyinkronkan memory antar device dengan prinsip local-first.

Contoh konfigurasi:

```yaml
cloud_memory:
  enabled: true
  provider: turso
  db_name_pattern: smara-{workspace}
  sync_interval_sec: 30
  conflict_policy: lww
  offline_mode: auto
  encrypt_at_rest: false
  max_rows_per_hour: 50000
  max_storage_mb: 8000
  embeddings_cloud: false
```

Token dan secret cloud tidak seharusnya disimpan langsung di YAML. Gunakan credential store/keyring atau environment variable sesuai konfigurasi deployment.

## Conflict policy

Saat dua device mengubah data yang sama, Smara dapat memakai policy:

| Policy | Fungsi |
| --- | --- |
| `lww` | Last-write-wins. Cocok untuk sinkron sederhana. |
| `manual` | Konflik ditahan untuk diselesaikan manual. |
| `archive-loser` | Versi kalah diarsipkan agar tidak hilang. |
| `merge-content` | Konten dicoba digabung saat aman. |

## Memory untuk dokumentasi

Untuk membuat docs Smara, memory dapat dipakai sebagai log keputusan:

```bash
smara memory save "Docs: VitePress dipilih sebagai framework utama" --tags docs,decision
smara memory save "Docs: Understand Anything dipakai nanti sebagai audit/gap analysis" --tags docs,planning
```

Lalu gunakan search saat update docs:

```bash
smara memory search "Docs:" --hybrid
```

## Best practice

- Simpan fakta jangka panjang, bukan noise percakapan sementara.
- Tulis memory dalam kalimat jelas dan spesifik.
- Pakai tags dan category sejak awal.
- Jangan simpan API key atau secret mentah.
- Backup sebelum migrasi, cleanup besar, atau perubahan schema.
- Untuk tim, dokumentasikan memory penting ke docs Markdown agar tidak hanya tersimpan di database lokal.
