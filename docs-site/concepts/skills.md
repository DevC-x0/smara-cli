# Skills

Skill adalah resep otomasi yang dapat dijalankan ulang oleh Smara. Skill cocok untuk workflow berulang seperti deploy, backup, monitoring, release, atau generate dokumentasi.

## Command umum

```bash
smara skill run deploy-backend
smara skill create deploy-backend --format json
smara skill install https://example.com/skills/deploy.json
smara skill search "deploy"
smara skill info deploy-backend
smara skill delete deploy-backend
```

## Registry

```bash
smara skill registry sync
smara skill publish deploy-backend
```

## Skill Tree & Analytics

```bash
smara skill tree
smara skill stats deploy-backend
smara skill analytics
smara skill refine deploy-backend
```

## Kapan membuat skill?

Buat skill jika workflow:

- Punya 3+ langkah tool call.
- Sering diulang.
- Punya parameter yang bisa diganti.
- Perlu standar tim.

Contoh kandidat skill:

- Build dan release CLI.
- Deploy static docs.
- Backup database VPS.
- Audit health service.
- Generate feature guide dari source code.
