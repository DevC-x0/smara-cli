# Golang Refactor Example

Contoh workflow refactor project Go dengan Smara.

## Prompt

```text
review package Go ini, usulkan refactor kecil yang aman, lalu jalankan go test setelah saya setuju
```

## Steps umum

1. Analisis struktur package.
2. Baca fungsi dan dependency utama.
3. Buat rencana refactor.
4. Edit file secara bertahap.
5. Jalankan formatter dan test.

## Commands

```bash
gofmt -w ./...
go test ./...
```

## Safety

Untuk refactor besar, minta Smara membuat plan dan rollback strategy sebelum mengubah file.
