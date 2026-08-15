# Policy Reference

Policy mengatur apakah tool atau operasi tertentu boleh dijalankan langsung, harus meminta approval, atau ditolak.

## Command

```bash
smara policy list
smara policy set <tool> <allow|ask|deny>
smara policy check <tool>
```

Contoh:

```bash
smara policy set ssh_exec ask
smara policy set delete_file ask
smara policy set run_command allow
smara policy check ssh_exec
```

## Mode policy

| Value | Arti | Cocok untuk |
|---|---|---|
| `allow` | Tool boleh dijalankan langsung | operasi read-only atau dev lokal aman |
| `ask` | Smara harus minta approval | edit file penting, deploy, restart service |
| `deny` | Tool tidak boleh dijalankan | operasi destructive/berisiko tinggi |

## Rekomendasi local development

```bash
smara policy set view_file allow
smara policy set grep_search allow
smara policy set run_command ask
smara policy set edit_file ask
smara policy set delete_file ask
```

## Rekomendasi production/server

```bash
smara policy set ssh_exec ask
smara policy set ssh_upload ask
smara policy set ssh_download ask
smara policy set delete_file deny
```

Gunakan `ask` untuk command yang dapat mengubah server, restart service, migrasi database, atau deploy.

## Hubungan dengan agent mode

- **Ask Mode**: default lebih konservatif, cocok untuk diskusi.
- **Plan Mode**: Smara menyusun rencana sebelum tool mutating/remote-write.
- **Rush Mode**: Smara lebih cepat eksekusi, tetapi policy tetap menjadi pagar.

## Best practice

- Set remote write ke `ask`.
- Set destructive tool ke `ask` atau `deny`.
- Review policy sebelum memberi akses ke workspace production.
- Simpan policy berbeda untuk environment berbeda jika workflow kamu membutuhkannya.
