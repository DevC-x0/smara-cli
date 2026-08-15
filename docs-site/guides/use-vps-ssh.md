# SSH Remote Control

Smara dapat mengelola VPS/server langsung dari agen dan CLI. Fitur ini cocok untuk deploy, monitoring, membaca log, upload/download file, dan operasi maintenance yang perlu dilakukan berulang.

## Prinsip aman

Untuk server production:

- gunakan SSH key, bukan password jika memungkinkan;
- simpan host dengan nama jelas seperti `prod`, `staging`, atau `docs`;
- jalankan command read-only dulu sebelum restart/delete;
- gunakan mode plan/approval untuk operasi berisiko;
- jangan hardcode secret di command atau skill;
- selalu verifikasi setelah perubahan.

## Tambah host

```bash
smara ssh add-host prod --host 192.168.1.10 --user ubuntu --key ~/.ssh/id_rsa
```

Contoh dengan port custom:

```bash
smara ssh add-host prod --host 203.0.113.10 --user ubuntu --port 2222 --key ~/.ssh/prod_ed25519
```

Lihat host:

```bash
smara ssh list
```

## Eksekusi command

```bash
smara ssh exec prod "docker ps -a"
smara ssh exec prod "systemctl status nginx --no-pager"
smara ssh exec prod "journalctl -u smara -n 100 --no-pager"
```

Untuk command multi-step, lebih aman tulis eksplisit:

```bash
smara ssh exec prod "cd /opt/app && git status --short && docker compose ps"
```

## Sesi interaktif

```bash
smara ssh connect prod
```

Gunakan ini jika perlu shell manual. Untuk workflow yang bisa diaudit, lebih baik pakai `ssh exec` atau skill agar command tercatat.

## Transfer file

Upload:

```bash
smara ssh upload prod ./local-file.txt /home/ubuntu/local-file.txt
```

Download:

```bash
smara ssh download prod /var/log/app.log ./logs/app.log
```

Contoh deploy artifact docs:

```bash
cd docs-site
npm run docs:build
smara ssh upload prod .vitepress/dist/index.html /var/www/docs/index.html
```

Untuk folder penuh, biasanya lebih praktis gunakan `rsync` via `ssh exec` atau buat archive terlebih dahulu.

## Key dan logs

```bash
smara ssh keygen --name deploy-key --type ed25519
smara ssh logs --limit 20
smara ssh transfer-logs --limit 20
```

Log membantu audit siapa/apa yang menjalankan command dan transfer.

## Agent tools

Saat `smara start`, agen bisa memakai tool SSH seperti:

- `ssh_exec` — menjalankan command remote;
- `ssh_view_file` — membaca file remote;
- `ssh_list_dir` — melihat folder remote;
- `ssh_upload` — upload file;
- `ssh_download` — download file.

Contoh prompt:

```text
Cek status nginx di vps prod, baca 100 log terakhir, lalu beri ringkasan tanpa restart dulu.
```

Untuk operasi production, minta agen membuat rencana dulu:

```text
Buat rencana restart service smara di prod, cek risiko, tunggu approval sebelum eksekusi.
```

## Workflow monitoring server

```bash
smara ssh exec prod "hostname && uptime"
smara ssh exec prod "df -h"
smara ssh exec prod "free -h"
smara ssh exec prod "systemctl --failed --no-pager"
smara ssh exec prod "journalctl -p err -n 50 --no-pager"
```

## Workflow deploy aplikasi

Contoh pattern aman:

1. Cek status repo/service.
2. Pull/update artifact.
3. Install/build jika perlu.
4. Restart service.
5. Verifikasi health check.
6. Baca log error terakhir.

```bash
smara ssh exec prod "cd /opt/smara && git status --short"
smara ssh exec prod "cd /opt/smara && git pull --ff-only"
smara ssh exec prod "cd /opt/smara && make build"
smara ssh exec prod "sudo systemctl restart smara"
smara ssh exec prod "systemctl status smara --no-pager"
```

## Jadikan skill

Jika workflow punya 3+ step dan sering diulang, simpan sebagai skill:

```json
{
  "name": "check-prod-health",
  "description": "Cek health dasar VPS production.",
  "params": [
    { "name": "host", "type": "string", "required": true, "default": "prod" }
  ],
  "steps": [
    { "tool": "ssh_exec", "args": { "host": "__PARAM__host", "command": "uptime" } },
    { "tool": "ssh_exec", "args": { "host": "__PARAM__host", "command": "df -h" } },
    { "tool": "ssh_exec", "args": { "host": "__PARAM__host", "command": "systemctl --failed --no-pager" } }
  ],
  "tags": ["ssh", "monitoring"]
}
```

## Troubleshooting

| Masalah | Solusi |
| --- | --- |
| Permission denied | Cek user, key path, permission `chmod 600`, dan authorized_keys. |
| Host tidak dikenal | Jalankan `smara ssh list`, pastikan nama host benar. |
| Command butuh sudo | Gunakan user yang punya sudo, hindari prompt password di automation. |
| Upload gagal | Cek path remote, permission folder, dan disk space. |
| Service tidak restart | Baca `systemctl status` dan `journalctl -u <service>`. |
