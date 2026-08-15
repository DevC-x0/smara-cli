# Release Checklist Example

Checklist ini membantu merilis Smara CLI secara konsisten.

## Pre-release

```bash
git status --short
go test ./...
cd web && npm run build
cd ../docs-site && npm run docs:build
cd .. && node scripts/audit-docs-cli.mjs
```

Pastikan:

- semua test lulus;
- web build lulus;
- docs build lulus;
- audit docs CLI tidak missing;
- `VERSION` sudah benar;
- release notes di `versions/RELEASE_vX.Y.Z.md` tersedia.

## Build assets

Contoh nama asset updater:

```text
smara-vX.Y.Z-linux-amd64.tar.gz
smara-vX.Y.Z-linux-arm64.tar.gz
smara-vX.Y.Z-darwin-amd64.tar.gz
smara-vX.Y.Z-darwin-arm64.tar.gz
smara-vX.Y.Z-windows-amd64.zip
SHA256SUMS.txt
```

## Checksums

```bash
cd dist
sha256sum * > SHA256SUMS.txt
cd ..
```

## Git tag

```bash
git tag vX.Y.Z
git push origin vX.Y.Z
```

## GitHub release

Upload:

- compressed binaries;
- `SHA256SUMS.txt`;
- release notes.

## Update smoke test

```bash
smara update --version vX.Y.Z --no-restart
smara doctor
smara --version
```

## Rollback

Jika ada masalah:

```bash
smara update --version vX.Y.(Z-1) --no-restart
```

Atau restore binary dari backup lokal.
