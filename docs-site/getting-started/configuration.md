# Configuration

Konfigurasi Smara disimpan sebagai file lokal user dan workspace. Jalankan:

```bash
smara config list
```

## Contoh config

```yaml
provider: anthropic
model: claude-3-5-sonnet-latest
ollama_host: http://localhost:11434
agent_max_iterations: 30

platforms:
  telegram:
    enabled: true
    token: YOUR_TOKEN
    allowed_users:
      - "12345678"
  discord:
    enabled: false
  whatsapp:
    enabled: false

skill_registries:
  - name: default
    url: https://raw.githubusercontent.com/gede-cahya/Smara-Skills/main/registry.json

cloud_memory:
  enabled: false
  conflict_policy: lww
```

## Provider

Gunakan command berikut untuk login dan memilih model:

```bash
smara login
smara provider list
smara provider select
```

## Workspace

Workspace memisahkan memori, sesi, dan konteks project.

```bash
smara workspace create client-x
smara workspace use client-x
smara workspace list
```

## Safety policy

Smara menyediakan policy CLI untuk mengatur allow/ask/deny pada tool automation.

```bash
smara policy
```

Gunakan mode `plan` untuk workflow yang memerlukan approval sebelum perubahan file, command shell, SSH, atau deployment.
