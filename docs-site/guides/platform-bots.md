# Platform Bots

Smara dapat dijalankan sebagai bot di platform messaging agar tim bisa memakai agent yang sama dari Telegram, Discord, atau WhatsApp.

```bash
smara serve
smara serve --platform telegram
smara serve --platform telegram,discord,whatsapp --mode plan
```

## Kapan dipakai

Gunakan platform bot ketika Smara perlu menjadi assistant tim:

- menjawab pertanyaan operasional singkat,
- menjalankan prompt agent dari chat,
- membantu monitoring server,
- mengakses memory/workspace bersama,
- memberi entrypoint non-terminal untuk tim.

Untuk operasi berisiko seperti deploy, restart service, atau command destructive, gunakan mode `plan` dan policy `ask`.

## Platform selection

Flag utama:

```bash
smara serve --platform telegram
smara serve --platform discord
smara serve --platform whatsapp
smara serve --platform telegram,discord,whatsapp
```

Jika `--platform` kosong, Smara mencoba menjalankan platform yang enabled di config/environment.

## Mode agent

```bash
smara serve --mode ask
smara serve --mode plan
smara serve --mode rush
```

Rekomendasi:

| Mode | Cocok untuk |
|---|---|
| `ask` | Bot publik/team dengan kontrol lebih aman. |
| `plan` | Operasi server, deploy, dan perubahan multi-step. |
| `rush` | Bot internal terpercaya untuk automation cepat. |

## Token dan config

Token dapat disimpan melalui config atau environment variable.

```bash
SMARA_TELEGRAM_TOKEN=bot123:AAH...
SMARA_DISCORD_TOKEN=MTIz...
```

Simpan secret di environment/server secret manager. Jangan commit token ke repository.

## MCP dan tools

Saat `smara serve` berjalan, Smara juga memuat provider LLM, memory store, dan MCP server yang enabled. Ini membuat bot bisa memakai tool eksternal yang sama dengan CLI.

Workflow aman:

1. Jalankan bot dengan `--mode plan`.
2. Set policy untuk command berisiko.
3. Batasi token hanya untuk channel/team terpercaya.
4. Monitor logs saat pertama kali deploy.

## Safety checklist

- Pakai token khusus bot, bukan token personal.
- Jangan expose bot publik dengan tool shell/SSH bebas.
- Gunakan allowlist command untuk production.
- Audit prompt yang meminta upload/download file.
- Pisahkan workspace untuk bot team dan local development.
