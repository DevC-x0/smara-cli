# MCP

MCP atau **Model Context Protocol** memungkinkan Smara memakai tool eksternal sebagai bagian dari workflow agen. Dengan MCP, Smara dapat terhubung ke browser automation, database inspector, docs fetcher, knowledge service, atau tool internal perusahaan.

Smara juga bisa berjalan sebagai MCP server sehingga editor/agent lain dapat memakai tool Smara.

## Mode penggunaan

Ada dua arah integrasi:

### 1. Smara sebagai MCP client

Smara membaca konfigurasi MCP, menghubungkan server, mengambil daftar tool, lalu membuat tool tersebut tersedia untuk agen.

Sumber konfigurasi yang dapat dimuat:

- konfigurasi Windsurf IDE;
- konfigurasi OpenCode;
- konfigurasi native Smara di `~/.smara/config.yaml`;
- remote MCP via endpoint HTTP/SSE/WebSocket sesuai transport yang didukung.

### 2. Smara sebagai MCP server

Smara dapat diekspose ke client lain:

```bash
smara mcp serve
```

Mode ini menjalankan server MCP stdio sehingga editor atau agent lain dapat memakai tool Smara seperti file tools, command tools, memory tools, dan workflow tools.

## Command

```bash
smara mcp list
smara mcp add filesystem --type local --command npx --args "-y,@modelcontextprotocol/server-filesystem,."
smara mcp add docs-api --type remote --url https://example.com/mcp
smara mcp remove filesystem
smara mcp serve
```

Untuk remote MCP dengan headers/env:

```bash
smara mcp add internal-docs \
  --type remote \
  --url https://mcp.example.com \
  --headers "Authorization=Bearer $TOKEN"
```

## Config native Smara

```yaml
mcp_servers:
  - name: filesystem
    type: local
    command: npx
    args:
      - -y
      - "@modelcontextprotocol/server-filesystem"
      - .
    env:
      NODE_ENV: production
    enabled: true

  - name: internal-docs
    type: remote
    url: https://mcp.example.com
    headers:
      Authorization: Bearer ${MCP_TOKEN}
    enabled: true
```

## Auto-discovery

Saat menjalankan supervisor/agent, Smara mencoba memuat konfigurasi dari beberapa tempat. Konfigurasi yang namanya sama akan dideduplicate agar server tidak terkoneksi dua kali.

Workflow umum:

1. Tambahkan MCP server ke config.
2. Jalankan `smara start` atau command yang membutuhkan supervisor.
3. Smara memuat dan menghubungkan MCP server yang enabled.
4. Smara membaca daftar tools.
5. Agen memakai tool sesuai policy dan izin.

## Kapan memakai MCP?

Gunakan MCP jika tool:

- sudah tersedia sebagai MCP server;
- terlalu spesifik untuk masuk core Smara;
- perlu dijalankan lintas editor/agent;
- punya lifecycle sendiri;
- membutuhkan integrasi enterprise seperti docs internal, database, atau observability.

Jangan pakai MCP jika workflow sederhana cukup dilakukan dengan tool built-in Smara seperti `run_command`, `read_file`, `ssh_exec`, atau `web_fetch`.

## Safety

MCP tools bisa punya dampak besar karena berasal dari server eksternal. Terapkan aturan berikut:

- hanya pasang MCP dari sumber tepercaya;
- cek tool list sebelum dipakai;
- gunakan environment variable untuk secret;
- nonaktifkan server yang tidak dipakai;
- untuk tool write/remote/destructive, jalankan dalam mode plan/approval;
- batasi scope server filesystem hanya ke folder yang diperlukan.

## MCP untuk dokumentasi

MCP berguna untuk docs workflow:

- Context7 untuk fetch dokumentasi library terbaru;
- browser MCP untuk cek halaman live docs;
- database MCP untuk inspeksi schema;
- internal knowledge MCP untuk menarik ADR/spec;
- Smara MCP server agar editor lain bisa menjalankan workflow docs Smara.

Contoh flow:

```text
Context7 MCP -> ambil docs library
Smara tools  -> scan repo
LLM          -> tulis Markdown
VitePress    -> build docs-site
```
