# Cloud Memory

Cloud Memory adalah mode sinkronisasi memory Smara lintas device/workspace dengan prinsip **local-first**. Database lokal tetap menjadi sumber kerja utama, sementara metadata cloud dipakai untuk replikasi, deduplication, dan conflict tracking.

## Kapan dipakai

Gunakan Cloud Memory jika kamu ingin:

- membawa konteks Smara ke beberapa mesin;
- menyimpan preferensi, project notes, dan docs decisions secara konsisten;
- memakai Smara di laptop + VPS/devbox;
- membuat knowledge base pribadi yang tetap tersedia lintas sesi;
- mengaudit atau memulihkan memory dari replica.

Untuk penggunaan satu mesin, SQLite lokal saja biasanya cukup.

## Model local-first

Smara tetap menyimpan data utama di:

```yaml
db_path: ~/.smara/memory.db
```

Saat Cloud Memory aktif, record memory diberi metadata tambahan seperti:

| Metadata | Fungsi |
|---|---|
| `cloud_id` | ID lintas device untuk dedup dan merge. |
| `origin_device_id` | Device asal write. |
| `content_hash` | Deteksi perubahan konten. |
| sync log | Melacak attempt, error, dan status delta sync. |
| conflict table | Menyimpan konflik divergent version. |

## Device ID

Setiap instalasi Smara memiliki device ID lokal, umumnya di area config Smara seperti:

```text
~/.smara/device-id
```

Device ID membantu membedakan write dari laptop, VPS, atau environment CI.

## Contoh config

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

Field penting:

| Field | Keterangan |
|---|---|
| `enabled` | Mengaktifkan cloud-aware memory store. |
| `provider` | Backend cloud memory. Saat ini diarahkan ke Turso/libSQL style setup. |
| `db_name_pattern` | Pola nama DB per workspace. |
| `sync_interval_sec` | Interval sync background. |
| `conflict_policy` | Cara menangani konflik. |
| `offline_mode` | Perilaku saat network/cloud gagal. |
| `max_rows_per_hour` | Guardrail agar sync tidak runaway. |
| `embeddings_cloud` | Apakah embedding ikut disimpan cloud. |

## Command umum

```bash
smara memory cloud login
smara memory cloud whoami
smara memory cloud status
smara memory cloud enable
smara memory cloud sync
smara memory cloud conflicts
smara memory cloud conflicts resolve <id>
smara memory cloud conflicts auto-resolve
smara memory cloud database list
smara memory cloud database info <nama-db>
smara memory cloud quota
smara memory cloud health
smara memory cloud token info
smara memory cloud encryption status
smara memory cloud encryption key-generate
smara memory cloud encryption key-delete
smara memory cloud disable
smara memory cloud logout
```

Danger zone:

```bash
smara memory cloud nuke
```

`nuke` bersifat irreversible untuk database cloud. Backup dan verifikasi workspace/provider sebelum menjalankannya.
- delta sync dicoba lagi nanti;
- error dicatat di sync log;
- konflik tidak boleh diam-diam menghapus data penting.

## Workflow aman enable cloud

1. Backup database lokal:

```bash
cp ~/.smara/memory.db ~/.smara/memory.backup.$(date +%Y%m%d-%H%M%S).db
```

2. Cek status memory:

```bash
smara memory list --limit 5
smara memory cloud status
```

3. Aktifkan Cloud Memory dan sync awal:

```bash
smara memory cloud enable
smara memory cloud sync
```

4. Cek konflik:

```bash
smara memory cloud conflicts
```

## Cloud Memory untuk docs

Gunakan memory sebagai log keputusan docs:

```text
ingat: Smara docs memakai VitePress, warna mengikuti desain lama, Understand Anything dipakai nanti untuk audit gap analysis.
```

Lalu cari kembali saat update docs:

```bash
smara memory search "Smara docs VitePress warna lama" --hybrid
```

## Security checklist

- Jangan simpan API key, private key SSH, token bot, atau credential mentah sebagai memory.
- Jika memakai provider cloud, pahami lokasi penyimpanan data.
- Gunakan workspace terpisah untuk project sensitif.
- Backup sebelum enable sync pertama kali.
- Aktifkan encryption-at-rest jika tersedia dan sesuai kebutuhan.
- Untuk tim, dokumentasikan keputusan penting ke Markdown, bukan hanya memory personal.
