# Sharing Reference

`smara share` mengelola metadata visibility untuk resource Smara seperti skill, workflow, dan memory.

```bash
smara share set skill deploy-project workspace
smara share set workflow release-checklist team --team core
smara share show skill deploy-project
smara share list
smara share list --type skill
```

## Resource type

Type yang didukung:

| Type | Keterangan |
|---|---|
| `skill` | Automation recipe yang bisa dijalankan ulang. |
| `workflow` | Orkestrasi multi-step/multi-role. |
| `memory` | Catatan/konteks yang bisa dibagikan berdasarkan metadata. |

Untuk `skill` dan `workflow`, Smara mengecek target benar-benar ada sebelum menyimpan metadata.

## Visibility

| Visibility | Arti |
|---|---|
| `private` | Hanya untuk user/workspace lokal. |
| `workspace` | Dibagikan dalam workspace aktif. |
| `team` | Dibagikan ke team tertentu. Gunakan `--team`. |

## Metadata path

`smara share list` menampilkan lokasi file metadata sharing. Lokasi aktual tergantung config/runtime Smara.

## Safety

- Jangan share skill yang menyimpan secret hardcoded.
- Review workflow sebelum visibility `team`.
- Untuk memory, cek apakah ada token, credential, atau data client.
- Pakai naming yang jelas: `deploy-staging`, `backup-prod-readonly`, `docs-generate`.
