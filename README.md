# MCP Shared Memory Server (Go) 🧠⚡

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go)](https://golang.org)
[![MCP Protocol](https://img.shields.io/badge/MCP-v1.0-blueviolet)](https://modelcontextprotocol.io)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

A ultra-fast, local-first Model Context Protocol (MCP) server written in pure Go. Enables seamless session handover and persistent cross-client knowledge sharing between AI Coding Agents (such as **Smara CLI**, **Antigravity**, **Claude Code**, **Cursor**, and **Windsurf**).

---

## 🌟 Key Features & Advantages

- **Zero-Dependency SQLite Engine**: Built using `modernc.org/sqlite` (CGO-free pure Go), ensuring painless cross-platform compilation and instant startup time (<5ms).
- **Concurrent-Safe Architecture**: WAL (Write-Ahead Logging) mode and busy timeout configuration guarantee zero database lock contention when multiple AI agents write/read simultaneously across different terminal sessions.
- **Cross-Session Auto Handover**: Seamlessly transition tasks, code context, and key decisions between different AI clients or sessions.
- **Persistent Knowledge Graph**: Store long-term project architecture, conventions, and configuration snippets accessible via fuzzy tag searches.
- **Local-First & Private**: Data stays 100% on your local machine (`~/.mcp-shared-memory/db.sqlite`).

---

## 🛠️ MCP Tools Reference

The server exposes 4 standard MCP tools:

### 1. `save_session_handover`
Saves current session progress, context, and remaining tasks to transfer to the next session.
- **Arguments**:
  - `project_name` *(string, required)*: Identifier for the workspace/project.
  - `summary` *(string, required)*: Concise recap of completed actions.
  - `remaining_tasks` *(array of strings)*: Action items for the next session.
  - `key_decisions` *(array of strings)*: Architectural or setup decisions made.
  - `code_context` *(string)*: Relevant file paths, signatures, or diff summaries.

### 2. `get_session_handover`
Retrieves the most recent session handover entry for a given project.
- **Arguments**:
  - `project_name` *(string, required)*: Identifier for the workspace/project.

### 3. `store_knowledge`
Stores long-term facts, conventions, environment configurations, or design decisions.
- **Arguments**:
  - `category` *(string, required)*: E.g., `architecture`, `config`, `credential_reference`, `convention`.
  - `title` *(string, required)*: Short title.
  - `content` *(string, required)*: Markdown or plain text body.
  - `tags` *(array of strings)*: List of searchable tags (e.g., `["db", "sqlite"]`).

### 4. `query_knowledge`
Queries knowledge nodes by keyword matching across titles, contents, and tags.
- **Arguments**:
  - `query` *(string)*: Search text.
  - `tag` *(string)*: Specific tag filter.
  - `limit` *(int, default=10)*: Maximum items returned.

---

## 🏷️ Tags & Categories Taxonomy Guide

Organizing memory with standardized tags ensures high retrieval accuracy across AI agents:

| Category | Example Tags | Typical Use Case |
| :--- | :--- | :--- |
| `architecture` | `backend`, `frontend`, `db`, `service-map` | System topology, DB schemas, component flows |
| `config` | `env`, `ports`, `nginx`, `docker` | Local dev setup, port bindings, config paths |
| `convention` | `formatting`, `git`, `linter`, `rules` | Team coding standards, branch rules |
| `decisions` | `adr`, `migration`, `tech-stack` | Architectural Decision Records (ADRs) |

---

## 🚀 How It Works (Multi-Agent Lifecycle)

```
+------------------+         +----------------------------+         +--------------------+
|  Claude Code     |         |   mcp-shared-memory        |         |  Antigravity /     |
|  (Session A)     |         |   SQLite Engine            |         |  Smara CLI         |
+--------+---------+         +-------------+--------------+         +---------+----------+
         |                                 |                                  |
         | --- 1. store_knowledge -------> | [Writes to ~/.mcp-shared-memory] |
         | --- 2. save_session_handover -> |                                  |
         |                                 |                                  |
         |                                 | <--- 3. get_session_handover --- |
         |                                 | <--- 4. query_knowledge -------- |
```

1. **Before Session Closes (Agent A)**: Calls `save_session_handover` to freeze state.
2. **On New Session Start (Agent B)**: Automatically calls `get_session_handover` to restore state and pick up where Agent A left off.
3. **During Execution**: Agents store and query reusable facts via `store_knowledge` and `query_knowledge`.

---

## 📦 Installation & Setup

### 1. Build Binary
```bash
go build -o bin/mcp-shared-memory-go ./cmd/server
```

### 2. Auto-Install Lifecycle Hooks
Run the automated installer to register MCP and lifecycle agent prompts for **Claude Code** and **Smara CLI**:
```bash
./install-hooks.sh
```

### 3. Client Manual Configurations

<details>
<summary><b>Claude Code (~/.claude.json)</b></summary>

```json
{
  "mcpServers": {
    "mcp-shared-memory": {
      "command": "/absolute/path/to/bin/mcp-shared-memory-go"
    }
  }
}
```
</details>

<details>
<summary><b>Smara CLI / Antigravity (~/.smara/mcp.json)</b></summary>

```json
{
  "mcpServers": {
    "mcp-shared-memory": {
      "command": "/absolute/path/to/bin/mcp-shared-memory-go"
    }
  }
}
```
</details>

---

## 🏷️ Latest Release Notes (v1.0.0)

### Version 1.0.0 (Initial Production Release)
- **Features**:
  - Implemented 4 core MCP tools (`save_session_handover`, `get_session_handover`, `store_knowledge`, `query_knowledge`).
  - Added WAL mode & busy handler lock mitigation for concurrent agent access.
  - Automated installer script `install-hooks.sh` for global agent system prompt integration.
- **Compatibility**: Standard MCP Protocol JSON-RPC over `stdio`.

---

## 📄 License

Distributed under the MIT License. See `LICENSE` for details.
