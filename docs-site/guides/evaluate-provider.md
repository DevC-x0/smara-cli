# Evaluate Provider

Smara menyediakan command evaluasi untuk menguji provider/model aktif terhadap suite sederhana. Ini berguna sebelum mengganti model default, menjalankan workflow penting, atau membandingkan kualitas provider.

```bash
smara eval run
smara eval run --json
smara eval run --file eval-suite.json
```

## Use cases

- Membandingkan OpenAI, OpenRouter, Anthropic, Ollama, atau custom provider.
- Mengecek provider baru sebelum dipakai di agent mode.
- Validasi regresi setelah update model.
- Membuat baseline latency/quality untuk docs dan release notes.

## Default suite

Jika `--file` tidak diberikan, Smara menjalankan default suite dari source internal eval.

Output terminal menampilkan:

- nama suite,
- jumlah pass/fail,
- latency per case,
- error atau missing expectation.

## JSON output

Gunakan `--json` untuk CI atau report otomatis.

```bash
smara eval run --json > eval-result.json
```

JSON output cocok digabungkan dengan workflow release:

1. build binary,
2. run eval,
3. run docs build,
4. generate release notes.

## Custom suite

```bash
smara eval run --file ./eval-suite.json
```

Gunakan custom suite untuk menguji prompt yang penting untuk workflow Smara, misalnya:

- membaca struktur repo,
- membuat rencana deploy,
- menjelaskan skill format,
- menghasilkan Markdown docs.

## Recommended workflow

```bash
smara provider set openrouter
smara provider set-model anthropic/claude-sonnet-4.5
smara eval run --json > reports/eval-openrouter.json

smara provider set anthropic
smara eval run --json > reports/eval-anthropic.json
```

Bandingkan hasilnya sebelum memilih provider default.
