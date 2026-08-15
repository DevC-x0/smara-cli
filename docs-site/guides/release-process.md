# Release Process

This guide documents the recommended release flow for Smara CLI maintainers.

## Goals

A good Smara release should include:

- reproducible version metadata
- Linux, macOS, and Windows assets
- SHA256 checksums
- release notes in `versions/`
- GitHub Release upload
- smoke tests for CLI, web, and docs
- safe update path via `smara update`

## Version sources

Common files involved in a release:

```text
VERSION
cmd/smara/version.go
versions/RELEASE_vX.Y.Z.md
dist/
build/release/
```

The Go variable in `cmd/smara/version.go` can be set at build time using linker flags. The `VERSION` file and docs package version should be kept aligned when preparing a public release.

## Pre-release checklist

```bash
git status --short
go test ./...
cd web && npm run build
cd ../docs-site && npm run docs:build
```

Recommended manual checks:

```bash
./smara version
./smara doctor
./smara provider list
./smara web --host 127.0.0.1 --port 7860
```

## Build assets

Expected release asset naming follows:

```text
smara-vX.Y.Z-linux-amd64.tar.gz
smara-vX.Y.Z-linux-arm64.tar.gz
smara-vX.Y.Z-darwin-amd64.tar.gz
smara-vX.Y.Z-darwin-arm64.tar.gz
smara-vX.Y.Z-windows-amd64.zip
```

The updater expects an asset matching the current OS/architecture. Keep names predictable:

```text
smara-{version}-{goos}-{goarch}.{tar.gz|zip}
```

## Generate checksums

```bash
cd dist
sha256sum smara-vX.Y.Z-* > SHA256SUMS.txt
```

Upload both archives and checksums to GitHub Release.

## Release notes

Create a release note file:

```text
versions/RELEASE_vX.Y.Z.md
```

Suggested structure:

```md
# Smara CLI vX.Y.Z

## Highlights

## Added

## Changed

## Fixed

## Upgrade Notes

## Checksums
```

Focus on user-visible changes, compatibility notes, and migration steps.

## Git tag and GitHub Release

```bash
git tag vX.Y.Z
git push origin vX.Y.Z
```

Then create or update the GitHub Release for `vX.Y.Z` and upload assets from `dist/`.

## Verify updater compatibility

`smara update` reads GitHub Releases and selects the asset for the current runtime:

```text
https://api.github.com/repos/gede-cahya/Smara-CLI/releases/latest
https://api.github.com/repos/gede-cahya/Smara-CLI/releases/tags/vX.Y.Z
```

Checks:

```bash
smara update --version X.Y.Z --no-restart
```

Use `--no-restart` when testing locally or when you do not want Smara to schedule a systemd restart.

## Production update behavior

On Linux, after replacing the binary, Smara tries to detect running systemd services that use the current binary and schedules a restart using `systemd-run`.

Operational guidance:

- use `--no-restart` for manual maintenance windows
- run with sufficient permissions if binary path requires root
- keep a backup binary for rollback
- verify service logs after update

## Rollback

If a release is bad:

1. Mark the GitHub Release as pre-release or remove broken assets.
2. Re-upload fixed assets or publish patch version.
3. On servers, reinstall previous version:

```bash
smara update --version X.Y.(Z-1) --no-restart
sudo systemctl restart smara
```

If the updater itself fails, manually download the archive from GitHub Releases, extract the binary, and replace the deployed binary path.

## Documentation release

Docs should be built before tagging:

```bash
cd docs-site
npm run docs:build
```

If docs are deployed separately, verify:

```bash
npm run docs:preview
```

## Minimal release command sequence

```bash
git status --short
go test ./...
cd web && npm run build
cd ../docs-site && npm run docs:build
cd ..
# build cross-platform assets into dist/
cd dist && sha256sum smara-vX.Y.Z-* > SHA256SUMS.txt && cd ..
git tag vX.Y.Z
git push origin vX.Y.Z
```

## Release safety checklist

- [ ] No accidental secrets in assets or notes.
- [ ] `VERSION`, release notes, and tag match.
- [ ] Archives use updater-compatible names.
- [ ] SHA256 checksums generated.
- [ ] `go test ./...` passes.
- [ ] `web` build passes.
- [ ] `docs-site` build passes.
- [ ] Update path tested with `--no-restart`.
