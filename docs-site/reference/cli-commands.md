# CLI Commands

Referensi ini disusun dari source `cmd/smara/*.go` dan README agar docs mengikuti fitur aktual Smara. Gunakan `smara <command> --help` untuk melihat flag paling lengkap di versi yang sedang terpasang.

## Core session

| Command | Fungsi |
|---|---|
| `smara start` | Mulai sesi TUI interaktif Smara. |
| `smara guide` | Panduan interaktif fitur Smara langsung di terminal. |
| `smara version` | Tampilkan versi binary. |
| `smara update [version]` | Update ke versi terbaru atau versi tertentu. |
| `smara doctor` | Diagnosis kesehatan instalasi/config Smara. |
| `smara repair` | Perbaikan otomatis untuk masalah umum. |
| `smara dashboard` | Dashboard monitoring real-time. |
| `smara analytics` | Analitik token, cost, model, request, prompt, dan skill. |
| `smara web` | Jalankan Smara Web Interface. |
| `smara claude-hook` | Hook integrasi Claude Code untuk menjalankan Smara sebagai command hook. |



```bash
smara login
smara provider list
smara provider select
smara provider set <name>
smara provider set-model <model>
smara provider test

smara config list
smara config get <key>
smara config set <key> <value>
```

Provider yang umum dipakai: Ollama/local, Anthropic, OpenAI, dan OpenRouter. API key disimpan lewat flow login/config lokal, bukan di source repository.

## Workspace dan project exploration

```bash
smara workspace list
smara workspace create <name>
smara workspace use <name>

smara explore [path]
```

Workspace memisahkan konteks proyek, memory, dan sesi. `explore` membantu membaca struktur project sebelum agen membuat rencana atau mengedit file.

## Memory

```bash
smara memory save <content>
smara memory list
smara memory get <id>
smara memory search <query> --hybrid
smara memory update <id>
smara memory history <id>
smara memory rollback <id> <version>
smara memory export backup.zip --format zip
smara memory import backup.zip
smara memory cleanup
```

### Memory Graph

```bash
smara memory link <source-id> <target-id> --relation refines --weight 0.8
smara memory unlink <link-id>
smara memory links <memory-id>
smara memory autolink --threshold 0.78 --top-k 5
smara memory graph
smara memory graph --export graph.json
```

Memory Graph menghubungkan catatan dengan relasi seperti `refines`, `supports`, atau `follows`, lalu bisa divisualkan di browser atau Smara Web.

## Cloud Memory

```bash
smara memory cloud login
smara memory cloud enable
smara memory cloud status
smara memory cloud workspaces
smara memory cloud conflicts
smara memory cloud resolve <id>
smara memory cloud provider
smara memory cloud database list
smara memory cloud quota
smara memory cloud health
smara memory cloud encryption status
```

Cloud Memory bersifat local-first: Smara tetap membaca/menulis local database, lalu sinkron ke provider seperti Turso/libSQL.

## Categories

```bash
smara category list
smara category create <nama>
smara category get <id>
smara category update <id>
smara category delete <id>
smara category stats <id>
```

Kategori membantu merapikan memory berdasarkan domain seperti coding, deployment, research, atau client project.

## Skills

```bash
smara skill list
smara skill run <name>
smara skill create <name> --format json
smara skill install <url>
smara skill search <query>
smara skill recommend <query>
smara skill suggest <query>
smara skill info <name>
smara skill publish <name>
smara skill delete <name>
smara skill registry sync
smara skill tree
smara skill stats <name>
smara skill runs [name]
smara skill analytics
smara skill refine <name>
smara skill lint [name]
smara skill validate [name]
smara skill export
smara skill import
```

## SSH, deploy, and 9drive

```bash
smara ssh add-host prod --host 192.168.1.1 --user ubuntu --key ~/.ssh/id_rsa
smara ssh list
smara ssh exec prod "docker ps"
smara ssh connect prod
smara ssh upload prod ./file.txt /home/ubuntu/file.txt
smara ssh download prod /var/log/app.log ./logs/app.log
smara ssh keygen --name deploy-key --type ed25519
smara ssh logs --limit 20
smara ssh transfer-logs --limit 20

smara deploy install prod
smara deploy status prod
smara deploy logs prod
smara deploy stop prod
smara deploy uninstall prod

smara 9drive upload photo.jpg
smara 9drive upload *.png
smara 9drive upload --api-key 9d_live_xxx backup.zip
```

Gunakan mode `plan` untuk operasi server berisiko agar agen menjelaskan rencana sebelum restart/deploy. `smara 9drive` menyediakan integrasi cloud storage 9drive untuk upload file/artifact; API key bisa diberikan via `--api-key` atau config `ninedrive_api_key`, dan endpoint bisa dioverride dengan `--base-url`.

## MCP

```bash
smara mcp list
smara mcp add <name>
smara mcp remove <name>
smara mcp serve
```

Smara dapat auto-discover MCP server dari konfigurasi IDE/agent lain dan juga dapat berjalan sebagai MCP server via stdio.

## Graphify — code knowledge graph

```bash
smara graphify init ./cmd --name smara-cmd
smara graphify query "auth flow" --name smara-cmd --depth 2
smara graphify path "A" "B" --name smara-cmd
smara graphify explain "NodeID" --name smara-cmd --depth 1
smara graphify export --name smara-cmd --format json
smara graphify export --name smara-cmd --format svg
smara graphify export --name smara-cmd --format graphml
smara graphify list
smara graphify delete smara-cmd
```

Graphify dipakai untuk workflow dokumentasi ala Understand Anything: scan repo, map fitur, cari gap docs, lalu generate Markdown.

## Automation, policy, schedule, sharing

```bash
smara workflow create <name>
smara workflow list
smara workflow run <name>
smara workflow show <name>
smara workflow delete <name>

smara run history
smara run show <id>
smara run replay <id>
smara run timeline <id>

smara schedule add <spec> <workflow>
smara schedule list
smara schedule remove <id>
smara schedule run-due
smara schedule daemon

smara policy list
smara policy set <tool> <allow|ask|deny>
smara policy check <tool>

smara share set <type> <name> <private|workspace|team>
smara share show <type> <name>
smara share list
```

## Evaluation, browser, image, voice, desktop

```bash
smara eval run
smara eval run --file eval-suite.json --json
smara browser run --url http://localhost:5173 --screenshot
smara browser e2e --spec browser-task.md
smara image "pastel green CLI docs hero" --out hero.png
smara voice transcribe --audio audio.wav
smara voice speak --text "Build selesai"
smara voice plan --text "buka browser dan cek docs"
smara desktop-agent --mode supervised
smara magic-pointer
```

Command ini melengkapi Smara sebagai agentic terminal: evaluasi provider, browser subagent, generate image, voice assistant, desktop bridge, dan Magic Pointer.

## Serve bot

```bash
smara serve
smara serve --platform telegram --mode plan
smara serve --platform telegram,discord,whatsapp --mode ask
```

Platform bot mendukung perintah seperti `/ask`, `/mode`, `/mcp`, dan `/clear` sesuai platform yang dikonfigurasi.
