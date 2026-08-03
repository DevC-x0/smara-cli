# mcp-shared-memory-go

MCP Cross-Session Memory Server berbasis Golang dan SQLite Pure Go (`modernc.org/sqlite`).

## Structure
- `cmd/server/main.go`: Entry point MCP stdio server.
- `pkg/db/sqlite.go`: SQLite connection & SQL query logic.
- `pkg/mcp/tools.go`: Handler 4 MCP Tools (`save_session_handover`, `get_session_handover`, `store_knowledge`, `query_knowledge`).

## Build & Test
```bash
go test -v ./...
go build -o bin/mcp-shared-memory-go ./cmd/server
```

## Configuration
Tambahkan ke `mcpServers` pada `.claude/mcp.json` atau config Antigravity:
```json
{
  "mcpServers": {
    "shared-memory": {
      "command": "/home/cahya/2026/mcp-bridge-memory/bin/mcp-shared-memory-go"
    }
  }
}
```
