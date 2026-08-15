# Cloud Memory Sync Example

Contoh workflow aman untuk mengaktifkan dan memverifikasi Cloud Memory.

## Goal

Sinkronisasi memory antar device tanpa mengorbankan prinsip local-first.

## Enable checklist

1. Backup memory lokal.
2. Pastikan device ID unik.
3. Login cloud memory.
4. Jalankan sync kecil.
5. Cek conflict dan quota.

## Commands

```bash
smara memory export smara-memory-backup.zip --format zip
smara memory cloud login
smara memory cloud whoami
smara memory cloud status
smara memory cloud sync
smara memory cloud conflicts
smara memory cloud quota
```

## Conflict handling

Jika ada konflik:

```bash
smara memory cloud conflicts list
smara memory cloud conflicts auto-resolve --strategy newest
```

Gunakan auto-resolve hanya jika kamu yakin strategi cocok. Untuk data penting, review manual dulu.

## Expected report

```text
Cloud account: user@example.com
Device: laptop-dev
Local memories: 1240
Uploaded: 8
Downloaded: 3
Conflicts: 0
Quota: 12% used
```

## Safety notes

- Jangan sync secret/API key sebagai memory biasa.
- Pakai kategori/retention untuk data sensitif.
- Backup sebelum conflict resolution massal.
