# Server Monitoring Example

Contoh workflow monitoring server dengan Smara dan SSH tools.

## Prompt

```text
cek status server vps-cahya: disk, memory, service nginx, dan log error terbaru. Jangan ubah apa pun dulu.
```

## Checks

- uptime
- disk usage
- memory usage
- service status
- recent logs
- open ports jika dibutuhkan

## Commands umum

```bash
df -h
free -h
systemctl status nginx --no-pager
journalctl -u nginx -n 80 --no-pager
```

## Skill opportunity

Jika workflow ini sering dipakai, simpan sebagai skill monitoring agar bisa dijalankan ulang dengan satu perintah.
