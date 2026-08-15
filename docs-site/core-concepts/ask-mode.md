# Ask Mode

Ask Mode adalah pola penggunaan Smara untuk membaca, menjelaskan, dan memberi saran tanpa melakukan perubahan langsung pada file atau server.

Gunakan mode ini ketika kamu ingin:

- memahami error atau log
- meminta review arsitektur
- membandingkan opsi teknis
- menyusun prompt, rencana, atau dokumentasi
- mengecek risiko sebelum eksekusi

## Karakteristik

- Tidak menjalankan aksi mutating kecuali diminta eksplisit.
- Cocok untuk eksplorasi dan tanya jawab.
- Bisa memakai tool read-only untuk memahami konteks, seperti membaca file, melihat struktur repo, atau mengambil dokumentasi.

## Contoh

```text
jelaskan struktur repo ini
```

```text
review rencana deploy ini, jangan eksekusi dulu
```

```text
apa risiko kalau saya ganti konfigurasi nginx ini?
```

::: tip
Untuk perubahan besar, gunakan Plan Mode agar Smara membuat rencana, meminta approval, lalu mengeksekusi bertahap.
:::
