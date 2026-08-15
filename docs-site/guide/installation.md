# Installation

Smara CLI dapat dipasang lewat installer resmi atau dibangun dari source.

## Prasyarat

- OS: Linux, macOS, atau Windows.
- Untuk build manual: Go 1.23+.
- Opsional: Ollama untuk model lokal.
- Opsional: Node.js/Wails untuk aplikasi desktop.

## Linux / macOS

```bash
curl -fsSL https://raw.githubusercontent.com/gede-cahya/Smara-CLI/main/install.sh | sh
```

Setelah selesai, cek versi:

```bash
smara version
```

## Windows PowerShell

```powershell
irm https://raw.githubusercontent.com/gede-cahya/Smara-CLI/main/install.ps1 | iex
```

## Build dari source

```bash
git clone https://github.com/gede-cahya/Smara-CLI.git
cd Smara-CLI
go build -o smara ./cmd/smara
./smara version
```

## Desktop app

Smara juga punya aplikasi desktop berbasis Wails.

```bash
cd smara-desktop
wails dev
# atau
wails build
```

## Next step

Lanjut ke [Quickstart](/getting-started/quickstart) untuk login provider dan menjalankan sesi pertama.
