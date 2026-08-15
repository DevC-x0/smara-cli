# Smara Desktop Rust MVP Scope

## Goal

Deliver a lightweight desktop MVP for Smara that provides a simple local UI for running Smara-like interactions, inspecting output, and preparing the integration path to the existing Smara CLI.

## MVP principles

- Lightweight first.
- Local-first by default.
- Reuse Smara CLI instead of rewriting core logic immediately.
- Keep UI minimal and fast.
- Make every external command auditable.

## In scope

### 1. Desktop shell

- Tauri v2 app skeleton.
- Rust backend command layer.
- Vanilla TypeScript frontend.
- Basic app layout.

### 2. Fake CLI bridge

- A Rust command that accepts a prompt string.
- Returns deterministic fake Smara output.
- Used to validate UI flow before real CLI integration.

### 3. Minimal UI

- Prompt input.
- Run button.
- Output panel.
- Basic status/error message area.

### 4. Developer documentation

- Local run instructions.
- Build instructions.
- Next steps for real CLI bridge.

## Out of scope for MVP

- Full Smara CLI rewrite in Rust.
- Multi-agent orchestration UI.
- Full skill/workflow marketplace.
- Cloud sync.
- Auto-update system.
- Production signing/notarization.
- Advanced terminal emulation.

## Acceptance criteria

The MVP foundation is accepted when:

- `smara-desktop-rust/` exists and contains a Tauri-oriented project skeleton.
- The frontend can call a Rust command named `run_fake_smara`.
- The fake command returns visible output for a prompt.
- The project has README instructions for development.
- Documentation exists for stack decision, architecture, and MVP scope.

## Next milestone after MVP foundation

Replace the fake bridge with a controlled real bridge to the existing `smara` CLI binary:

- Detect binary path.
- Execute commands safely.
- Stream stdout/stderr.
- Add approval boundaries for mutating commands.
- Store local run history.
