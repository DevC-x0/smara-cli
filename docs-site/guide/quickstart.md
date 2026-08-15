# Quickstart

Panduan ini membantu menjalankan Smara pertama kali.

## 1. Login provider

```bash
smara login
```

Ikuti prompt untuk menyimpan API key provider. Smara mendukung beberapa provider seperti Ollama, Anthropic, OpenAI, dan OpenRouter.

## 2. Pilih model

```bash
smara provider select
```

Untuk model lokal, pastikan Ollama berjalan:

```bash
ollama serve
ollama pull llama3.1
```

## 3. Mulai TUI

```bash
smara start
```

Di TUI, kamu bisa bertanya, meminta Smara menganalisis project, atau menjalankan workflow.

## 4. Pilih mode agen

Smara punya tiga mode utama:

- `ask`: tanya jawab cepat tanpa tool mutating.
- `rush`: eksekusi cepat untuk tugas yang sudah jelas.
- `plan`: menyusun rencana dan meminta approval sebelum aksi berisiko.

Rekomendasi awal: gunakan `plan` untuk perubahan code, server, deployment, atau file penting.

## 5. Coba command dasar

```bash
smara guide
smara config list
smara provider list
smara dashboard --once
```

## 6. Workflow contoh

```text
Tolong analisis repo ini dan buat rencana refactor struktur command.
```

Smara akan membaca context, membuat rencana, lalu menunggu approval sebelum menjalankan perubahan.
