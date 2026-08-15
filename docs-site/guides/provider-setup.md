# Provider Setup

Smara mendukung beberapa provider LLM: **Ollama**, **OpenAI**, **OpenRouter**, **Anthropic**, dan **custom OpenAI-compatible endpoint**. Provider aktif menentukan model yang dipakai oleh ask mode, agent mode, workflow, skills, dan sebagian tool multimodal.

## Cek provider tersedia

```bash
smara provider list
# alias
smara provider
```

Output menampilkan:

- provider yang tersedia;
- status API key/login;
- provider aktif;
- daftar model default;
- apakah provider local atau cloud.

## Login provider cloud

```bash
smara login --provider openai
smara login --provider openrouter
smara login --provider anthropic
```

Jika tidak memberi `--provider`, Smara menampilkan status dan memberi petunjuk login.

::: warning
Jangan commit API key ke repo. Simpan lewat mekanisme login/config lokal atau environment variable sesuai deployment.
:::

## Ganti provider aktif

```bash
smara provider set openai
smara provider set openrouter
smara provider set anthropic
smara provider set ollama
smara provider set custom
```

Jika provider membutuhkan API key dan belum login, command akan menolak dan memberi instruksi login.

## Ganti model

```bash
smara provider set-model gpt-4.1-mini
```

Command ini mengubah `model` aktif dan juga menyimpan model khusus provider bila provider mendukung field khusus seperti:

```yaml
openai_model: gpt-4.1-mini
openrouter_model: openai/gpt-4.1-mini
anthropic_model: claude-3-5-sonnet-latest
custom_model: llama3.1
```

## Test koneksi

```bash
smara provider test
```

Smara akan mengirim pesan sederhana:

```text
Reply with 'OK' if you can read this.
```

Jika sukses, output menampilkan model dan token. Jika gagal, periksa API key, base URL, provider aktif, dan koneksi internet.

## Provider selector TUI

```bash
smara provider select
```

Gunakan untuk memilih provider/model secara interaktif jika tersedia di environment terminal.

## Ollama local

Ollama cocok untuk workflow lokal/private.

```bash
ollama serve
ollama pull llama3.1
smara provider set ollama
smara provider set-model llama3.1
smara provider test
```

Config umum:

```yaml
provider: ollama
model: llama3.1
ollama_host: http://localhost:11434
```

## Custom OpenAI-compatible provider

Gunakan `custom` untuk endpoint seperti LocalAI, LiteLLM, vLLM, atau gateway internal.

```bash
smara login --provider custom
# atau
smara login --custom
```

Field yang biasanya diminta:

```yaml
provider: custom
custom_provider_name: local-ai
custom_base_url: http://localhost:8080/v1
custom_api_key: sk-local-or-empty
custom_model: llama3.1
```

Setelah setup:

```bash
smara provider set custom
smara provider test
```

## OpenRouter

OpenRouter berguna untuk membandingkan banyak model melalui satu endpoint.

```bash
smara login --provider openrouter
smara provider set openrouter
smara provider set-model openai/gpt-4.1-mini
smara provider test
```

## Provider untuk docs generation

Untuk membuat dokumentasi panjang:

- pilih model dengan konteks cukup besar;
- gunakan `--budget` saat mengambil graph/query output;
- build docs setelah update Markdown;
- jangan kirim secret atau private customer data ke provider cloud tanpa izin.

Workflow rekomendasi:

```bash
smara provider list
smara provider set openrouter
smara provider set-model openai/gpt-4.1-mini
smara provider test
cd docs-site && npm run docs:build
```

## Troubleshooting cepat

| Masalah | Cek |
|---|---|
| `provider tidak dikenali` | Gunakan `smara provider list`. |
| `memerlukan API key` | Jalankan `smara login --provider <name>`. |
| custom endpoint gagal | Cek `custom_base_url`, path `/v1`, dan API key. |
| Ollama timeout | Pastikan `ollama serve` aktif dan model sudah di-pull. |
| model tidak valid | Gunakan nama model sesuai provider. |
