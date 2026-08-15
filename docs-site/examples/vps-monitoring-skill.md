# VPS Monitoring Skill Example

Contoh ini mengubah prompt monitoring server yang sering diulang menjadi skill Smara yang reusable.

## Kapan dipakai

Gunakan skill ini jika kamu sering meminta Smara untuk mengecek:

- uptime dan load;
- disk usage;
- memory usage;
- status service penting;
- log error terbaru.

## Prompt manual awal

```text
Cek status server vps-cahya: uptime, disk, memory, service nginx, dan log error terbaru. Jangan ubah apa pun.
```

Jika hasilnya bagus dan sering dipakai, simpan sebagai skill.

## Skill JSON

```json
{
  "name": "monitor-vps-basic",
  "description": "Read-only health check untuk VPS: uptime, disk, memory, service, dan log terbaru.",
  "tags": ["vps", "monitoring", "ssh"],
  "params": [
    { "name": "host", "type": "string", "required": true, "default": "vps-cahya" },
    { "name": "service", "type": "string", "required": false, "default": "nginx" }
  ],
  "steps": [
    { "tool": "ssh_exec", "args": { "host": "__PARAM__host", "command": "uptime && df -h && free -h" } },
    { "tool": "ssh_exec", "args": { "host": "__PARAM__host", "command": "systemctl status __PARAM__service --no-pager || true" } },
    { "tool": "ssh_exec", "args": { "host": "__PARAM__host", "command": "journalctl -u __PARAM__service -n 80 --no-pager || true" } }
  ]
}
```

## Jalankan

```bash
smara skill run monitor-vps-basic --param host=vps-cahya --param service=nginx
```

Atau dari chat:

```text
Jalankan skill monitor-vps-basic untuk host vps-cahya service nginx.
```

## Expected output

Ringkasan ideal:

```text
Status: OK / warning / critical
Disk: root 42% used
Memory: 1.2G / 4G used
Service nginx: active
Recent errors: none / found N lines
Recommended action: no action / investigate logs
```

## Safety checklist

- Skill ini read-only: tidak menjalankan `restart`, `rm`, `apt upgrade`, atau edit file.
- Untuk production, gunakan mode `plan` sebelum remediation.
- Jangan simpan token/password di skill JSON.
- Batasi log output agar tidak memenuhi context.
