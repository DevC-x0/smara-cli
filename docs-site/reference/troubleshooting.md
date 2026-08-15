# Troubleshooting

Panduan cepat untuk masalah umum saat memakai Smara CLI, docs-site, provider, memory, SSH, MCP, dan automation.

## `smara` command tidak ditemukan

Pastikan binary ada di PATH.

```bash
which smara
smara version
```

Jika build manual:

```bash
go build -o smara ./cmd/smara
sudo mv smara /usr/local/bin/
```

Jika binary ada di folder lokal, tambahkan ke PATH:

```bash
export PATH="$PWD:$PATH"
```

## Provider gagal login

Cek API key dan provider aktif:

```bash
smara login
smara provider list
smara provider select
```

Untuk provider cloud:

```bash
smara login --provider openai
smara login --provider openrouter
smara login --provider anthropic
smara provider test
```

## Provider aktif tapi model gagal

```bash
smara provider list
smara provider set-model <model>
smara provider test
```

Pastikan nama model sesuai provider. OpenRouter biasanya memakai format seperti:

```text
openai/gpt-4.1-mini
anthropic/claude-3.5-sonnet
```

## Ollama tidak merespons

```bash
ollama serve
curl http://localhost:11434/api/tags
ollama pull llama3.1
```

Pastikan config:

```yaml
provider: ollama
model: llama3.1
ollama_host: http://localhost:11434
```

Lalu test:

```bash
smara provider set ollama
smara provider test
```

## Custom provider gagal

Cek base URL dan model:

```yaml
provider: custom
custom_base_url: http://localhost:8080/v1
custom_model: llama3.1
```

Lalu:

```bash
smara provider set custom
smara provider test
```

Jika endpoint OpenAI-compatible, pastikan route `/v1/chat/completions` tersedia.

## Cloud Memory bermasalah

```bash
smara memory cloud status
smara memory cloud sync
smara memory cloud conflicts
```

Langkah aman:

1. Backup `~/.smara/memory.db`.
2. Nonaktifkan cloud sementara jika sync loop/error.
3. Cek conflict log.
4. Enable ulang untuk bootstrap replica.

```bash
cp ~/.smara/memory.db ~/.smara/memory.backup.$(date +%Y%m%d-%H%M%S).db
```

## Memory search tidak menemukan hasil

Coba longgarkan query dan gunakan hybrid search jika tersedia:

```bash
smara memory list --limit 20
smara memory search "deploy docs" --hybrid
smara memory search "VitePress"
```

Pastikan memory disimpan di workspace yang benar.

## Memory Graph kosong

Memory graph butuh memory dan link.

```bash
smara memory list --limit 10
smara memory autolink --threshold 0.78 --top-k 5
smara memory graph
```

Jika belum ada cukup data, simpan beberapa memory dulu atau buat link manual.

## SSH gagal connect

Cek host, user, port, dan key:

```bash
ssh -i ~/.ssh/id_rsa ubuntu@HOST
smara ssh list
smara ssh logs --limit 20
```

Checklist:

- security group/firewall membuka port SSH;
- key path benar dan permission aman (`chmod 600 key`);
- username benar (`ubuntu`, `root`, `deploy`, dll);
- host tersimpan mengarah ke IP terbaru.

## MCP server tidak muncul

```bash
smara mcp list
smara mcp discover
```

Untuk local MCP:

- command tersedia di PATH;
- args benar;
- environment variable lengkap;
- server tidak hang saat start.

Untuk remote MCP:

- URL benar;
- header auth benar;
- endpoint bisa diakses dari mesin Smara.

## Skill gagal dijalankan

```bash
smara skill list
smara skill info <name>
smara skill run <name>
```

Cek:

- nama tool di step benar;
- parameter wajib terisi;
- placeholder `__PARAM__nama` sesuai definisi params;
- command mutating punya approval jika policy membutuhkannya.

## Browser automation gagal

```bash
smara browser run "buka halaman docs dan cek title"
```

Cek:

- browser/headless dependency tersedia;
- target URL bisa diakses;
- login/session tidak expired;
- jangan jalankan aksi destruktif di production tanpa approval.

## Build docs gagal

Untuk docs VitePress:

```bash
cd docs-site
npm install
npm run docs:build
```

Jika package lock lama bermasalah:

```bash
rm -rf node_modules
npm install
npm run docs:build
```

Cek juga:

- link internal sidebar valid;
- file Markdown ada;
- frontmatter YAML valid;
- custom HTML tertutup dengan benar.

## Graphify gagal parse

```bash
smara graphify init . --name smara
```

Cek:

- jalankan dari root module Go;
- repo tidak memiliki file generated/vendor terlalu besar;
- permission file cukup;
- pakai path eksplisit jika perlu.

```bash
smara graphify init /path/to/repo --name smara
smara graphify query "memory" --name smara --depth 2
```

## Update gagal

```bash
smara update
smara update --version <tag>
```

Jika service sedang berjalan, gunakan checklist production:

1. Backup config dan database.
2. Stop service jika perlu.
3. Update binary.
4. Restart service.
5. Cek `smara doctor` atau command health terkait.

## Mode aman saat ragu

Jika task berisiko, minta Smara untuk planning dulu:

```text
buat rencana dulu, jangan eksekusi command mutating sebelum saya approve
```

Untuk production, prefer policy `ask` atau `deny` pada tool berisiko.
