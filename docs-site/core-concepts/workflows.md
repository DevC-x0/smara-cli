# Workflows

Workflow adalah orchestration layer untuk tugas multi-step yang lebih besar dari satu prompt atau satu skill. Di Smara, workflow cocok untuk pekerjaan yang butuh beberapa fase, role berbeda, shared state, dan resume.

## Workflow vs skill

| Aspek | Skill | Workflow |
|---|---|---|
| Tujuan | Recipe singkat yang sering diulang | Proses multi-step/multi-role |
| Bentuk umum | Urutan tool calls dengan parameter | Blueprint, roles, shared state, runner |
| Cocok untuk | backup, deploy sederhana, cek server | build product, audit repo, docs generation besar |
| Resume | Terbatas | Lebih cocok untuk task panjang |

Gunakan skill jika alurnya pendek dan stabil. Gunakan workflow jika tugas punya banyak fase dan perlu koordinasi.

## Command dasar

```bash
smara workflow create docs-refresh
smara workflow list
smara workflow show docs-refresh
smara workflow run docs-refresh
smara workflow delete docs-refresh
smara workflow import docs-refresh
```

## Konsep internal

Source workflow berada di `internal/agent/workflow/`. Beberapa konsep penting:

- **Blueprint**: rancangan tugas dan step utama.
- **Roles**: spesialis seperti frontend, backend, devops, QA, content strategist, legal researcher, data engineer, binary analyst, dan lain-lain.
- **Shared state**: konteks yang dipakai lintas step/role.
- **Runner**: eksekutor workflow.
- **Resume**: mekanisme melanjutkan workflow panjang.
- **Skill map**: pemetaan skill yang bisa membantu step tertentu.

## Contoh use case

### Generate dokumentasi besar

```text
Buat workflow untuk scan repo, identifikasi fitur, update docs VitePress, lalu build test.
```

Fase yang mungkin:

1. Repo analysis.
2. Feature map.
3. Gap report.
4. Markdown generation.
5. Review and polish.
6. Build verification.

### Release preparation

```text
Buat workflow release: test, build cross-platform, generate release notes, upload asset GitHub release.
```

### Production audit

```text
Buat workflow audit VPS: cek disk, service, docker, logs, backup config, lalu buat laporan.
```

## Best practice

- Mulai dengan plan read-only.
- Pecah workflow menjadi fase kecil.
- Gunakan skill untuk step yang sudah stabil.
- Simpan output penting sebagai artifact atau memory.
- Selalu punya verification step.
- Untuk remote/server, gunakan policy `ask` atau mode `plan`.

## Kapan jangan memakai workflow

Jangan pakai workflow jika:

- tugas hanya satu command sederhana,
- tidak perlu shared state,
- tidak perlu resume,
- hasilnya cukup dengan satu skill pendek.

Untuk kasus seperti itu, skill atau prompt langsung lebih cepat.
