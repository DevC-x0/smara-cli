# Web Interface

Smara Web Interface menjalankan UI lokal untuk chat real-time, memory, workspace, config, dan dashboard monitoring.

```bash
smara web
smara web --port 3000
smara web --host 0.0.0.0 --port 8080 --auth-token "$SMARA_WEB_TOKEN"
```

## Default behavior

Secara default:

- host: `127.0.0.1`,
- port: `8080`,
- mode: `ask`,
- membuka browser otomatis kecuali `--no-open`.

## Flags penting

| Flag | Fungsi |
|---|---|
| `--port` | Port HTTP server. |
| `--host` | Host listen. Pakai `0.0.0.0` hanya jika perlu akses network. |
| `--no-open` | Jangan buka browser otomatis. |
| `--mode` | Mode agent default: `ask`, `rush`, atau `plan`. |
| `--auth-token` | Token akses remote via header Authorization atau query `?token=`. |
| `--desktop-agent` | URL desktop-agent untuk pairing remote desktop. |
| `--desktop-token` | Token desktop-agent. |

## Remote access safety

Jika expose web interface ke network:

1. Gunakan `--auth-token`.
2. Pasang reverse proxy HTTPS.
3. Batasi firewall/IP.
4. Pakai mode `plan` untuk produksi.
5. Jangan expose endpoint local tanpa auth.

## Desktop Agent pairing

Smara Web bisa dipair dengan desktop-agent lokal:

```bash
smara desktop-agent --token "$DESKTOP_TOKEN"
smara web --desktop-agent http://127.0.0.1:8765 --desktop-token "$DESKTOP_TOKEN"
```

Gunakan pairing hanya di mesin terpercaya karena desktop-agent dapat mengontrol mouse, keyboard, clipboard, dan aplikasi.
