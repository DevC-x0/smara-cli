# Docs Site Maintenance

This example shows a repeatable workflow for maintaining the VitePress documentation site from the Smara source tree.

## Objective

Keep docs accurate while avoiding accidental changes to runtime code.

## Maintenance loop

```text
scan source -> identify docs gaps -> edit Markdown -> build -> review -> commit
```

## 1. Check workspace status

```bash
git status --short
```

If there are unrelated modified runtime files, treat them as protected unless the task explicitly includes them.

## 2. Scan CLI commands

```bash
grep -R "Use:" cmd/smara internal -n | sort
```

Use the output to update:

```text
docs-site/reference/cli-commands.md
docs-site/reference/tool-list.md
docs-site/reference/docs-gap-report.md
```

## 3. Scan web API routes

```bash
grep -R "HandleFunc" internal/web -n
```

Use the output to update:

```text
docs-site/reference/web-api.md
docs-site/guides/web-interface.md
```

## 4. Scan feature modules

Useful source directories:

```text
internal/memory
internal/skill
internal/mcp
internal/graphify
internal/web
internal/ssh
internal/eval
internal/scheduler
internal/runlog
internal/sharing
```

Map each module to one docs page:

| Source | Docs |
|---|---|
| `internal/memory` | `core-concepts/memory.md`, `guides/cloud-memory.md` |
| `internal/skill` | `core-concepts/skills.md`, `reference/skill-format.md` |
| `internal/mcp` | `core-concepts/mcp.md` |
| `internal/graphify` | `reference/graphify.md`, `guides/knowledge-graph.md` |
| `internal/web` | `reference/web-api.md`, `guides/web-interface.md` |
| `internal/ssh` | `guides/use-vps-ssh.md` |
| `internal/eval` | `guides/evaluate-provider.md` |
| `internal/scheduler` | `guides/scheduled-automation.md` |

## 5. Build docs

```bash
cd docs-site
npm run docs:build
```

Optional preview:

```bash
npm run docs:preview
```

## 6. Audit docs vs CLI

Run the audit script:

```bash
node scripts/audit-docs-cli.mjs
```

It scans Cobra command definitions and checks whether command names appear in docs pages.

## 7. Suggested commit shape

Keep docs-only changes separate from runtime fixes:

```bash
git add docs-site scripts/audit-docs-cli.mjs
git commit -m "docs: update Smara documentation coverage"
```

For runtime fixes, use a separate commit:

```bash
git add internal/web web/src
git commit -m "fix(web): improve session handling"
```

## Agent prompt example

```text
Scan cmd/smara and internal/web read-only, update only docs-site and scripts/audit-docs-cli.mjs, then run docs build. Do not modify runtime files.
```

## Checklist

- [ ] `git status --short` reviewed.
- [ ] CLI command reference updated.
- [ ] Web API reference updated if routes changed.
- [ ] Gap report updated.
- [ ] `npm run docs:build` passes.
- [ ] Runtime files untouched unless explicitly requested.


## CI / local audit shortcut

Docs package menyediakan shortcut:

```bash
cd docs-site
npm run docs:audit
npm run docs:check
```

`docs:check` menjalankan build VitePress lalu audit coverage CLI. CI GitHub Actions memakai alur yang sama via `.github/workflows/docs-audit.yml`.
