# SSH Remote Control

Smara dapat mengelola VPS/server langsung dari agen dan CLI.

## Tambah host

```bash
smara ssh add-host prod --host 192.168.1.10 --user ubuntu --key ~/.ssh/id_rsa
```

## Eksekusi command

```bash
smara ssh exec prod "docker ps -a"
smara ssh exec prod "systemctl status nginx --no-pager"
```

## Sesi interaktif

```bash
smara ssh connect prod
```

## Transfer file

```bash
smara ssh upload prod ./local-file.txt /home/ubuntu/local-file.txt
smara ssh download prod /var/log/app.log ./logs/app.log
```

## Key dan logs

```bash
smara ssh keygen --name deploy-key --type ed25519
smara ssh logs --limit 20
smara ssh transfer-logs --limit 20
```

## Agent tools

Saat `smara start`, agen bisa mendapatkan tool seperti:

- `ssh_exec`
- `ssh_view_file`
- `ssh_list_dir`
- `ssh_upload`
- `ssh_download`

Gunakan mode `plan` untuk operasi production. Smara akan membuat rencana, menunggu approval, lalu mengeksekusi bertahap.
