# Web API Reference

Smara Web exposes a local HTTP API used by the React web interface. The API is intended for the bundled UI first, but it is also useful for local automation, debugging, and lightweight integrations.

> Safety: bind to `127.0.0.1` for local use. If you expose Smara Web remotely, enable an auth token and put it behind a trusted reverse proxy/TLS boundary.

## Start the server

```bash
smara web --host 127.0.0.1 --port 7860
# or
smara web --host 0.0.0.0 --port 7860 --auth-token "$SMARA_WEB_TOKEN"
```

Base URL examples:

```text
http://127.0.0.1:7860
https://smara.example.com
```

## Authentication

If an auth token is configured, send it with each request:

```bash
curl -H "Authorization: Bearer $SMARA_WEB_TOKEN" http://127.0.0.1:7860/api/status
```

## Common response pattern

Most endpoints return JSON. Mutating endpoints generally use `POST`, `PUT`, or `DELETE` and may return either a JSON object or an empty success response.

Frontend helpers should handle:

- JSON response bodies
- empty `204 No Content` responses
- non-2xx responses as errors

## Core endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/status` | `GET` | Runtime status, provider, mode, workspace. |
| `/api/chat` | `POST` | Send a prompt/message to the agent runtime. |
| `/api/mode` | `GET/POST` | Read or switch assistant mode. |
| `/api/config` | `GET/POST` | Read/update web-visible config keys. |
| `/api/metrics` | `GET` | Usage analytics and token/cost summaries. |
| `/ws` | WebSocket | Streaming/runtime events used by the UI. |

## Filesystem endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/fs/cwd` | `GET` | Current workspace directory. |
| `/api/fs/list` | `GET` | List files/directories for browser panels. |

## Memory endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/memories` | `GET/POST` | List or create memories. Supports workspace/limit query in UI flows. |
| `/api/memories/search` | `POST` | Search memory store. |
| `/api/memories/graph` | `GET` | Memory graph nodes/edges. |
| `/api/memories/links` | `GET/DELETE` | Inspect or delete links between memories. |
| `/api/memories/autolink` | `POST` | Build similarity/wikilink-based memory links. |

Example:

```bash
curl -X POST http://127.0.0.1:7860/api/memories/search \
  -H 'Content-Type: application/json' \
  -d '{"query":"release process","limit":20}'
```

## Workspace endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/workspaces` | `GET` | List workspaces and active workspace. |
| `/api/workspaces/switch` | `POST` | Switch active workspace. |
| `/api/workspaces/create` | `POST` | Create a workspace. |
| `/api/categories` | `GET/POST` | Memory category management. |

## Skill endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/skills` | `GET` | List installed skills. |
| `/api/skills/run` | `POST` | Run a skill. |
| `/api/skills/import` | `POST` | Import a skill. |
| `/api/skills/bundled` | `GET` | List bundled skills. |
| `/api/skills/install-bundled` | `POST` | Install bundled skill. |
| `/api/skills/tree` | `GET` | Skill hierarchy tree. |
| `/api/skills/stats` | `GET` | Skill usage stats. |
| `/api/skills/timeline` | `GET` | Skill execution timeline. |
| `/api/skills/analytics` | `GET` | Skill analytics. |
| `/api/skills/refine` | `POST` | Refine/improve a skill. |
| `/api/skills/dependencies` | `GET` | Skill dependency graph. |
| `/api/skills/export` | `GET` | Export skills/tree. |
| `/api/skills/import-tree` | `POST` | Import skill tree. |

## Workflow endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/blueprint/generate` | `POST` | Generate workflow/agent blueprint. |
| `/api/blueprint/execute` | `POST` | Execute generated blueprint. |
| `/api/custom-workflow/list` | `GET` | List custom workflows. |
| `/api/custom-workflow/get` | `GET` | Get custom workflow detail. |
| `/api/custom-workflow/save` | `POST` | Save workflow. |
| `/api/custom-workflow/delete` | `DELETE/POST` | Delete workflow. |
| `/api/custom-workflow/run` | `POST` | Run workflow. |
| `/api/custom-workflow/import` | `POST` | Import workflow. |

## Graphify endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/graph/init` | `POST` | Initialize a source-code graph. |
| `/api/graph/list` | `GET` | List graphs. |
| `/api/graph/get` | `GET` | Read graph metadata. |
| `/api/graph/query` | `POST` | Query graph by keyword. |
| `/api/graph/nodes` | `GET` | List graph nodes. |
| `/api/graph/neighbors` | `GET` | Get node neighborhood. |
| `/api/graph/path` | `GET/POST` | Find path between nodes. |
| `/api/graph/data` | `GET` | Full graph payload for visualization. |
| `/api/graph/export` | `GET` | Export graph data. |

## Artifacts and media endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/clipboard/upload` | `POST` | Upload clipboard content. |
| `/api/attachments/upload` | `POST` | Upload chat attachments. |
| `/api/generated-image` | `GET` | Serve generated image artifact. |
| `/api/local-image` | `GET` | Serve local image file safely. |
| `/api/browser-artifact` | `GET` | Serve browser automation artifact. |

## Multi-session chat endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/web-sessions` | `GET/POST` | List/create web sessions. Query: `include_archived`, `limit`. |
| `/api/web-sessions/{id}` | `GET/PUT/DELETE` | Get/update/delete session. Query: `limit` for history. |
| `/api/web-sessions/{id}/cancel` | `POST` | Cancel a running session. |
| `/api/web-sessions/{id}/archive` | `POST` | Archive/unarchive session depending payload. |

Notes:

- list calls should use a small history preview limit
- detail calls can load a larger history limit
- responses include `total_history` and `history_limit` when history is truncated

## Voice, avatar, and remote desktop endpoints

| Endpoint | Method | Purpose |
|---|---:|---|
| `/api/voice/settings` | `GET/POST` | Voice settings. |
| `/api/voice/command` | `POST` | Voice command endpoint. |
| `/api/voice/speak` | `POST` | Text-to-speech. |
| `/api/avatar/state` | `GET` | Avatar runtime state. |
| `/api/avatar/event` | `POST` | Avatar event. |
| `/api/avatar/model` | `GET/POST` | Avatar model settings. |
| `/api/remote-desktop/devices` | `GET` | List remote desktop devices. |
| `/api/remote-desktop/devices/{id}` | `GET/DELETE` | Device detail/removal. |
| `/api/remote-desktop/proxy` | `GET/POST` | Proxy remote desktop data. |
| `/api/remote-desktop/screenshot` | `GET` | Fetch screenshot. |


## Payload details for priority endpoints

### `POST /api/chat`

Send a message to the active web runtime. The UI uses this endpoint for non-streaming chat submission; streaming/status updates are delivered through `/ws` and session polling.

Request shape:

```json
{
  "message": "summarize this repo",
  "mode": "plan",
  "session_id": "optional-existing-session",
  "attachments": [
    { "name": "screenshot.png", "path": "/tmp/screenshot.png", "type": "image/png" }
  ]
}
```

Typical response:

```json
{
  "session_id": "web-20260521-abc123",
  "message": "accepted",
  "status": "running"
}
```

Use `/api/web-sessions/{id}?limit=...` to fetch the resulting history.

### `GET /api/web-sessions?include_archived=true&limit=5`

Returns a lightweight session list. Use a small `limit` for sidebar previews.

```json
{
  "sessions": [
    {
      "id": "web-20260521-abc123",
      "title": "Fix chat session race",
      "status": "completed",
      "archived": false,
      "history": [{ "role": "user", "content": "..." }],
      "total_history": 42,
      "history_limit": 5
    }
  ]
}
```

### `GET /api/web-sessions/{id}?limit=80`

Returns one session with a larger but still bounded history window.

```json
{
  "id": "web-20260521-abc123",
  "title": "Fix chat session race",
  "status": "completed",
  "history": [],
  "total_history": 42,
  "history_limit": 80
}
```

### `POST /api/memories/search`

Search local memory. Keep queries specific and use `limit` to avoid large UI payloads.

```json
{
  "query": "release checklist",
  "workspace": "smara",
  "limit": 20
}
```

Response:

```json
{
  "results": [
    { "id": "mem_123", "content": "...", "category": "project", "score": 0.82 }
  ]
}
```

### `POST /api/graph/query`

Query a Graphify graph. Use `depth` to control neighborhood size.

```json
{
  "graph": "smara",
  "query": "web session handlers",
  "depth": 2
}
```

Response shape depends on graph backend, but generally includes matching nodes, edges, and optional explanation text.

### `POST /api/skills/run`

Run an installed skill. Parameters map to skill placeholders.

```json
{
  "name": "build-smara",
  "params": {
    "target": "local"
  }
}
```

Response:

```json
{
  "status": "ok",
  "output": "...",
  "steps": 3
}
```

## Practical checks

```bash
curl http://127.0.0.1:7860/api/status
curl http://127.0.0.1:7860/api/workspaces
curl 'http://127.0.0.1:7860/api/web-sessions?limit=5'
```

## Stability notes

The Web API currently follows the web UI implementation. For external integrations, prefer:

1. CLI commands for stable automation.
2. MCP tools for tool-level integrations.
3. Web API only when you specifically need the local web runtime state.
