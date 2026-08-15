# Scheduled Automation

Schedule memungkinkan workflow dijalankan berdasarkan waktu. Ini cocok untuk monitoring, backup, cleanup, atau laporan berkala.

## Command

```bash
smara schedule add [spec] [workflow]
smara schedule list
smara schedule remove [id]
smara schedule run-due
smara schedule daemon
```

## Contoh

```bash
smara schedule add "daily at 09:00" morning-check
smara schedule add "every 30 minutes" server-health-check
smara schedule list
```

Jalankan due jobs secara manual:

```bash
smara schedule run-due
```

`run-due` berguna untuk cron/systemd timer karena hanya menjalankan job yang sudah jatuh tempo, lalu selesai.

Atau jalankan daemon:

```bash
smara schedule daemon
```

`daemon` cocok untuk proses long-running di server/devbox. Pastikan log dipantau dan workflow yang dijalankan aman untuk mode otomatis.
smara schedule daemon
```

`daemon` cocok untuk proses long-running di server/devbox. Pastikan log dipantau dan workflow yang dijalankan aman untuk mode otomatis.
- backup database,
- generate laporan token/cost,
- cleanup memory lama,
- refresh docs gap report,
- monitor endpoint aplikasi.

## Contoh workflow monitoring

```text
Workflow: server-health-check
Steps:
1. SSH ke host production.
2. Cek disk usage.
3. Cek memory dan load.
4. Cek service penting.
5. Ringkas status.
6. Jika error, buat alert.
```

## Safety

- Jangan jadwalkan destructive command tanpa guard.
- Untuk server production, jadwalkan read-only check atau backup yang sudah teruji.
- Simpan log hasil schedule.
- Gunakan policy `ask` untuk operasi yang tidak boleh otomatis.
- Test workflow manual sebelum dijadwalkan.
