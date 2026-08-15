# Cloud Memory

Cloud Memory membuat database memori Smara tetap local-first sekaligus tersinkronisasi ke Turso/libSQL untuk multi-device.

## Quick start

```bash
smara memory cloud login
smara memory cloud enable
smara memory cloud status
```

## Multi-device

Di device kedua:

```bash
smara memory cloud login
smara memory cloud enable
```

Smara akan memakai database cloud workspace yang sama bila tersedia.

## Workspace isolation

```bash
smara workspace create client-x
smara memory cloud workspaces
```

Untuk workspace lokal saja:

```bash
smara workspace create client-x --local-only
```

## Headless / CI

```bash
export SMARA_CLOUD_TOKEN=...
export SMARA_CLOUD_ORG=...
export SMARA_CLOUD_REGION=sin
smara memory cloud login --headless
smara memory cloud enable
```

## Conflict policy

Konfigurasi:

```yaml
cloud_memory:
  conflict_policy: lww
```

Pilihan:

- `lww`: last-write-wins deterministik.
- `manual`: konflik ditahan untuk review.
- `archive-loser`: versi kalah diarsipkan.
- `merge-content`: gabungkan konten catatan bebas.

## Troubleshooting cepat

- Keyring tidak tersedia: Smara fallback ke `~/.smara/cloud-creds.json` mode `0600`.
- Replica bermasalah: nonaktifkan cloud, backup `~/.smara/memory.db`, lalu enable ulang.
- Kuota penuh: cek `status`, cleanup data, atau naikkan kuota provider.
