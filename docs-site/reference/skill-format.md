# Skill Format

Skill adalah resep otomasi Smara yang berisi metadata dan urutan tool-call. Format utama adalah JSON, dan beberapa workflow juga dapat dibuat dari Markdown skill.

Skill disimpan sebagai file lokal dan dapat di-install dari bundled skills, registry, marketplace, atau URL langsung.

## Struktur JSON

```json
{
  "name": "deploy-docs-site",
  "description": "Build dan deploy docs-site ke server.",
  "version": 1,
  "author": "team-smara",
  "source_url": "https://example.com/skills/deploy-docs-site.json",
  "tags": ["docs", "deploy"],
  "params": [
    {
      "name": "host",
      "type": "string",
      "description": "Nama host SSH tujuan.",
      "required": true,
      "default": "prod"
    }
  ],
  "category_path": ["docs", "deploy"],
  "parent_id": "deploy-base",
  "dependencies": ["build-smara-docs"],
  "steps": [
    {
      "tool": "run_command",
      "args": {
        "command": "cd docs-site && npm run docs:build"
      }
    },
    {
      "tool": "ssh_exec",
      "args": {
        "host": "__PARAM__host",
        "command": "systemctl status nginx --no-pager"
      }
    }
  ]
}
```

## Field reference

| Field | Required | Fungsi |
| --- | --- | --- |
| `name` | Ya | Nama skill. Gunakan kebab-case. |
| `description` | Ya | Ringkasan apa yang dilakukan dan kapan dipakai. |
| `steps` | Ya | Urutan tool-call yang dijalankan. Minimal satu step. |
| `version` | Tidak | Versi skill. Naik saat skill direfine/update. |
| `tags` | Tidak | Kategori pendek untuk search/filter. |
| `author` | Tidak | Pembuat skill. |
| `source_url` | Tidak | URL asal skill untuk update/install. |
| `params` | Tidak | Parameter runtime. |
| `parent_id` | Tidak | Parent di skill tree. |
| `category_path` | Tidak | Jalur kategori bertingkat. |
| `dependencies` | Tidak | Skill lain yang dibutuhkan. |
| `lineage` | Otomatis | Riwayat versi lama setelah refinement. |

## Params

Parameter membuat skill reusable.

```json
{
  "name": "service",
  "type": "string",
  "description": "Nama systemd service.",
  "required": true,
  "default": "smara"
}
```

Tipe yang umum:

- `string`
- `number`
- `boolean`

Parameter dipakai di step args dengan placeholder:

```json
"command": "systemctl status __PARAM__service --no-pager"
```

Jalankan dengan argumen runtime:

```bash
smara skill run check-service --args '{"host":"prod","service":"nginx"}'
```

## Steps

Setiap step terdiri dari:

```json
{
  "tool": "run_command",
  "args": {
    "command": "go test ./..."
  }
}
```

`tool` harus sesuai tool yang tersedia di Smara/supervisor/MCP. Contoh:

- `run_command`
- `read_file`
- `write_file`
- `edit_file`
- `grep_search`
- `ssh_exec`
- `ssh_upload`
- `web_fetch`
- `graphify_init`
- `graphify_query`

## Placeholder substitution

Smara mengganti string `__PARAM__nama` secara rekursif di args.

Contoh:

```json
{
  "params": [
    { "name": "host", "type": "string", "required": true },
    { "name": "path", "type": "string", "required": true }
  ],
  "steps": [
    {
      "tool": "ssh_exec",
      "args": {
        "host": "__PARAM__host",
        "command": "ls -la __PARAM__path"
      }
    }
  ]
}
```

## Lineage dan refinement

Saat skill direfine, versi lama dapat disimpan di `lineage` agar histori tidak hilang.

Lineage menyimpan:

- versi sebelumnya;
- deskripsi lama;
- tags lama;
- jumlah step;
- waktu refinement;
- sumber refinement seperti `auto`, `manual`, atau `feedback`.

## Markdown skill

Selain JSON, command create mendukung format Markdown:

```bash
smara skill create my-skill --format md
```

Gunakan Markdown untuk skill yang ingin mudah dibaca manusia, lalu biarkan Smara parse menjadi struktur skill.

## Validasi minimum

Skill valid jika:

- `name` tidak kosong;
- `steps` tidak kosong;
- setiap step punya `tool`;
- args sesuai schema tool;
- placeholder parameter punya definisi param yang jelas.

## Best practice

- Buat nama spesifik: `deploy-docs-site`, bukan `deploy`.
- Jangan hardcode API key, token, password, atau private path user lain.
- Tambahkan step verifikasi di akhir.
- Gunakan parameter untuk host, branch, service, dan path.
- Pisahkan workflow destructive dari workflow read-only.
- Simpan skill production di registry yang terkontrol.
- Tambahkan dependency jika skill membutuhkan skill lain.

## Anti-pattern

Hindari:

```json
{
  "name": "do-everything",
  "steps": [
    { "tool": "run_command", "args": { "command": "curl unknown.sh | bash" } }
  ]
}
```

Lebih aman:

1. download script;
2. review isi;
3. jalankan hanya setelah dipercaya;
4. verifikasi hasil.
