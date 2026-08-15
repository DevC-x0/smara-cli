# Config Reference

Smara menyimpan konfigurasi utama di:

```text
~/.smara/config.yaml
```

Konfigurasi dibaca melalui Viper dan dapat diubah dari CLI:

```bash
smara config list
smara config get provider
smara config set provider anthropic
smara config set model claude-sonnet-4-20250514
```

> Jangan commit file config yang berisi token/API key ke repository.

## Provider dan model

Smara mendukung beberapa provider LLM. Field umum:

```yaml
provider: custom
model: deepseek-v4-pro
ollama_host: http://localhost:11434
```

Provider-specific:

```yaml
openai_api_key: ""
openai_model: gpt-4o
openai_base_url: ""

openrouter_api_key: ""
openrouter_model: anthropic/claude-sonnet-4

anthropic_api_key: ""
anthropic_model: claude-sonnet-4-20250514

custom_provider_name: CLIProxyAPI
custom_api_key: your-api-key-1
custom_base_url: http://localhost:8317/v1
custom_model: deepseek-v4-pro
```

Rekomendasi:

- gunakan environment variable/keyring untuk secret jika memungkinkan;
- pastikan `model` sinkron dengan provider aktif;
- untuk local model, set `ollama_host` dengan benar.

## Agent runtime

```yaml
agent_max_iterations: 80
agent_request_timeout_sec: 3600
platform_prompt_timeout: 600
verbose: false
```

| Field | Fungsi |
| --- | --- |
| `agent_max_iterations` | Batas loop think/tool/observe per task. |
| `agent_request_timeout_sec` | Timeout wall-clock untuk satu turn web/TUI. |
| `platform_prompt_timeout` | Timeout prompt dari Telegram/Discord/WhatsApp. |
| `verbose` | Output debug lebih detail. |

Iteration budget mencegah loop tool berulang, tetapi masih bisa diperpanjang untuk workflow kompleks jika agen meminta budget tambahan.

## Paths dan workspace

```yaml
sync_dir: ~/.smara/sync
db_path: ~/.smara/memory.db
active_workspace: default
image_output_dir: ~/.smara/images
```

| Field | Fungsi |
| --- | --- |
| `sync_dir` | Folder sinkronisasi lokal. |
| `db_path` | SQLite database memory. |
| `active_workspace` | Workspace aktif. |
| `image_output_dir` | Folder output image generation. |

## Image generation

```yaml
image_model: gpt-image-2
image_output_dir: ~/.smara/images
```

Gunakan ini saat Smara menjalankan tool image generation untuk logo, ilustrasi, mockup, atau asset docs.

## MCP servers

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

Field MCP:

| Field | Fungsi |
| --- | --- |
| `name` | Nama unik server. |
| `type` | `local` atau `remote`. |
| `command` | Command untuk local MCP. |
| `args` | Argumen command. |
| `url` | Endpoint remote MCP. |
| `headers` | Header HTTP remote MCP. |
| `env` | Environment variables local MCP. |
| `enabled` | Aktif/nonaktif. |

Smara juga punya field MCP internal:

```yaml
smara_mcp_enabled: false
smara_mcp_command: smara
smara_mcp_args:
  - mcp
  - serve
smara_mcp_api_key: ""
```

## Platforms

Smara dapat menerima prompt dari platform chat.

```yaml
platforms:
  telegram:
    enabled: true
    token: YOUR_TOKEN
    allowed_users:
      - "12345678"
    blocked_users: []
    owner_id: "12345678"
    sensitive_keywords:
      - password
      - token
    sensitive_deny_message: "Request ditolak karena mengandung data sensitif."
    rate_limit: 10
    rate_burst: 3

  discord:
    enabled: false
    token: YOUR_TOKEN
    guild_ids: []
    allowed_roles: []
    rate_limit: 10
    rate_burst: 3

  whatsapp:
    enabled: false
    session_dir: ~/.smara/wa-session
    allowed_numbers: []
    rate_limit: 10
    rate_burst: 3

  max_response_length: 4000
  typing_indicator: true
  log_conversations: false
```

## Skill registry dan auto skill

```yaml
skill_registries:
  - name: smara-official
    url: https://raw.githubusercontent.com/gede-cahya/smara-skills/main/skill-registry.json

auto_skill_detect: true
auto_skill_threshold: 3
```

| Field | Fungsi |
| --- | --- |
| `skill_registries` | Daftar registry marketplace skill. |
| `auto_skill_detect` | Deteksi workflow berulang secara otomatis. |
| `auto_skill_threshold` | Minimal kemunculan pola sebelum dicapture. |

## Cloud Memory

```yaml
cloud_memory:
  enabled: false
  provider: turso
  db_name_pattern: smara-{workspace}
  sync_interval_sec: 30
  conflict_policy: lww
  offline_mode: auto
  encrypt_at_rest: false
  max_rows_per_hour: 50000
  max_storage_mb: 8000
  embeddings_cloud: false
  sync_tables:
    - memories
    - memory_links
    - memory_versions
    - categories
    - workspaces
```

Conflict policy:

- `lww`
- `manual`
- `archive-loser`
- `merge-content`

## Minimal config example

```yaml
provider: custom
model: deepseek-v4-pro
custom_base_url: http://localhost:8317/v1
custom_api_key: ${SMARA_CUSTOM_API_KEY}
db_path: ~/.smara/memory.db
active_workspace: default
agent_max_iterations: 80
```

## Production safety checklist

- Jangan simpan token mentah di repo.
- Pakai `allowed_users`/`allowed_numbers` untuk platform bots.
- Batasi MCP filesystem scope.
- Untuk server production, gunakan policy approval/mode plan.
- Backup `~/.smara/config.yaml` dan `~/.smara/memory.db` sebelum migrasi.
- Naikkan timeout hanya jika workflow memang membutuhkan.
