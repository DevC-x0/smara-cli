# Skills

Skill adalah resep otomasi Smara yang dapat dijalankan ulang. Skill mengubah workflow multi-step menjadi satu command atau satu instruksi agen yang konsisten.

Contoh use case:

- build dan release CLI;
- deploy docs-site;
- backup database VPS;
- audit health service;
- generate dokumentasi dari source code;
- workflow Context7/MCP;
- refactor dan validasi project Go.

## Cara kerja

Sebuah skill berisi metadata dan urutan tool-call:

```text
nama + deskripsi + parameter + steps + tags + dependencies
```

Saat dijalankan, Smara:

1. membaca definisi skill;
2. mengisi parameter runtime dan default;
3. mengganti placeholder seperti `__PARAM__host`;
4. menjalankan steps secara berurutan;
5. menampilkan hasil tiap step;
6. menyimpan lineage/refinement jika skill diperbarui.

## Command umum

```bash
smara skill list
smara skill run deploy-backend
smara skill run deploy-backend --args '{"host":"prod"}'
smara skill create deploy-backend --format json
smara skill create docs-audit --format md
smara skill install context7-docs
smara skill install https://example.com/skills/deploy.json
smara skill info deploy-backend
smara skill delete deploy-backend
```

## Search dan install

`skill install` dapat mencari beberapa sumber:

1. bundled skills bawaan Smara;
2. Context7/library-style registry;
3. marketplace registry dari config;
4. URL langsung.

```bash
smara skill search "deploy"
smara skill install graphify-init
smara skill install context7-docs --as docs-context
smara skill install https://example.com/my-skill.json --overwrite
```

## Registry

Registry memudahkan distribusi skill ke tim atau publik.

```bash
smara skill registry list
smara skill registry sync
smara skill publish deploy-backend
```

Konfigurasi registry:

```yaml
skill_registries:
  - name: smara-official
    url: https://raw.githubusercontent.com/gede-cahya/smara-skills/main/skill-registry.json
```

## Skill Tree & Analytics

Skill bisa punya relasi parent/dependency sehingga membentuk skill tree.

```bash
smara skill tree
smara skill stats deploy-backend
smara skill analytics
smara skill refine deploy-backend
smara skill mock --prefix mock- --clean
```
```
Fitur ini membantu:

- melihat workflow yang saling bergantung;
- menemukan skill lanjutan yang bisa dibuat;
- memperbaiki skill berdasarkan histori eksekusi;
- menjaga lineage ketika skill direfine.

## Auto skill capture

Smara dapat mendeteksi pola tool-call yang berulang dan menyarankan atau membuat skill otomatis.

Konfigurasi terkait:

```yaml
auto_skill_detect: true
auto_skill_threshold: 3
```

Rekomendasi: aktifkan untuk workflow development, tapi tetap review hasil skill sebelum dipakai di production.

## Kapan membuat skill?

Buat skill jika workflow:

- punya 3+ langkah tool call;
- sering diulang;
- punya parameter yang bisa diganti;
- perlu standar tim;
- berisiko jika dilakukan manual berulang;
- punya verifikasi akhir yang jelas.

Contoh kandidat skill:

- `deploy-vitepress-docs`
- `audit-vps-service`
- `backup-postgres-vps`
- `generate-release-notes`
- `scan-repo-and-update-docs`

## Contoh skill docs-site

```json
{
  "name": "build-smara-docs",
  "description": "Build VitePress docs-site Smara dan verifikasi output static.",
  "version": 1,
  "tags": ["docs", "vitepress"],
  "steps": [
    {
      "tool": "run_command",
      "args": {
        "command": "cd docs-site && npm run docs:build"
      }
    }
  ]
}
```

## Safety

Skill dapat menjalankan command lokal, SSH, upload file, atau edit file. Karena itu:

- hindari hardcoded secret;
- gunakan parameter untuk host/path;
- tambahkan step verifikasi;
- pisahkan destructive operation ke skill khusus;
- untuk production, jalankan dalam mode plan/approval jika tersedia;
- baca `skill info` sebelum menjalankan skill dari URL eksternal.
