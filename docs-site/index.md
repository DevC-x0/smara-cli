---
layout: home

hero:
  name: Smara CLI
  text: Autonomous Multi-Agent Terminal
  tagline: Terminal AI berbasis Go untuk merencanakan, mengeksekusi, mengingat, dan mengotomasi workflow developer dengan safety gate.
  image:
    src: /logo.svg
    alt: Smara CLI
  actions:
    - theme: brand
      text: Mulai Cepat
      link: /getting-started/quickstart
    - theme: alt
      text: Generate Docs
      link: /guides/generate-docs
    - theme: alt
      text: Knowledge Graph
      link: /guides/knowledge-graph

features:
  - icon: 🧠
    title: Smart Memory
    details: "Hybrid search, kategorisasi, versioning, cloud sync, dan memory graph yang saling terhubung."
  - icon: 🧩
    title: Skill Ecosystem
    details: "Reusable automation skills dengan skill tree, dependency graph, analytics, dan refinement."
  - icon: 🌐
    title: MCP Native
    details: "Auto-discovery tool dari Windsurf, OpenCode, config Smara-native, dan remote MCP."
  - icon: 🖥️
    title: SSH Remote Control
    details: "Kelola VPS/server dari percakapan: exec, view file, list dir, upload, download, dan logs."
  - icon: 🕸️
    title: Graphify
    details: "Bangun knowledge graph dari codebase Go, query natural language, export JSON/SVG/GraphML/Neo4j."
  - icon: 🛡️
    title: Plan-first Safety
    details: "Mode plan dengan approval gate sebelum aksi mutating atau berisiko."
---

<div class="hero-card bg-grid">
  <strong>Smara</strong> berasal dari Sanskerta <em>smṛti</em> — ingatan. Smara dirancang sebagai terminal AI yang tidak hanya menjawab, tetapi juga mengingat, merencanakan, mengeksekusi, dan mengotomasi workflow developer secara aman.
</div>

<div class="badge-row">
  <span>Go</span>
  <span>Multi-Agent</span>
  <span>MCP</span>
  <span>Skills</span>
  <span>Cloud Memory</span>
  <span>SSH</span>
  <span>Graphify</span>
  <span>VitePress Docs</span>
</div>

## Documentation System

Smara doc site memakai konsep gabungan: **Understand Anything untuk analisis dan knowledge graph**, **VitePress untuk publishing**, dan **Smara/LLM untuk generate serta refine Markdown**.

<div class="smara-flow">
  <div class="smara-card"><strong>Analyze</strong><p>Scan repo Smara dan pahami struktur fitur, command, package, dan workflow.</p></div>
  <div class="smara-card"><strong>Map</strong><p>Bangun peta fitur dan knowledge graph agar docs mengikuti arsitektur aktual.</p></div>
  <div class="smara-card"><strong>Generate</strong><p>Buat draft Markdown untuk guide, reference, skill, dan onboarding.</p></div>
  <div class="smara-card"><strong>Refine</strong><p>Review isi, cek gap dokumentasi, lalu polish bahasa dan contoh command.</p></div>
  <div class="smara-card"><strong>Publish</strong><p>Build static site dengan VitePress dan deploy ke hosting pilihan.</p></div>
</div>

<div class="smara-orbit glow-border" aria-label="Smara documentation map">
  <div class="orbit-core">
    <strong>Smara</strong>
    <span>agent memory graph</span>
  </div>
  <a class="orbit-node orbit-node-a" href="/core-concepts/memory">Memory</a>
  <a class="orbit-node orbit-node-b" href="/core-concepts/skills">Skills</a>
  <a class="orbit-node orbit-node-c" href="/reference/graphify">Graphify</a>
  <a class="orbit-node orbit-node-d" href="/core-concepts/mcp">MCP</a>
  <a class="orbit-node orbit-node-e" href="/guides/provider-setup">Provider</a>
  <a class="orbit-node orbit-node-f" href="/guides/use-vps-ssh">SSH</a>
  <svg class="orbit-lines" viewBox="0 0 760 360" role="presentation" aria-hidden="true">
    <path d="M380 180 C230 70 120 120 105 180 C125 255 250 286 380 180Z"></path>
    <path d="M380 180 C525 60 650 105 660 180 C640 255 505 292 380 180Z"></path>
    <path d="M380 180 C300 40 465 36 380 180 C300 315 470 320 380 180Z"></path>
    <line x1="380" y1="180" x2="132" y2="88"></line>
    <line x1="380" y1="180" x2="632" y2="88"></line>
    <line x1="380" y1="180" x2="632" y2="272"></line>
    <line x1="380" y1="180" x2="132" y2="272"></line>
  </svg>
</div>


## Interactive docs graph

<DocsGraph />

## Quick install

::: code-group

```bash [Linux / macOS]
curl -fsSL https://raw.githubusercontent.com/gede-cahya/Smara-CLI/main/install.sh | sh
```

```powershell [Windows PowerShell]
irm https://raw.githubusercontent.com/gede-cahya/Smara-CLI/main/install.ps1 | iex
```

:::

## Start here

- Baru memakai Smara? Baca [Quickstart](/getting-started/quickstart).
- Ingin setup model? Baca [Provider Setup](/guides/provider-setup).
- Ingin memahami mode kerja agen? Baca [Agent Mode](/core-concepts/agent-mode).
- Mau kelola server dari chat? Baca [Use VPS / SSH](/guides/use-vps-ssh).
- Mau generate dokumentasi dari repo? Baca [Generate Docs](/guides/generate-docs).
- Mau dokumentasi knowledge graph? Baca [Knowledge Graph](/guides/knowledge-graph).
