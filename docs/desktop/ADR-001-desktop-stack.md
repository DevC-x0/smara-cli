# ADR-001: Smara Desktop Rust Stack

## Status

Accepted for MVP.

## Context

Smara needs a lightweight desktop application that can provide a friendly UI while reusing the existing Smara CLI capabilities. The first desktop version should be fast to start, small enough to distribute, and safe to evolve without rewriting the entire Smara core.

## Decision

Use **Rust + Tauri v2** for the desktop shell.

Initial stack:

- Rust for desktop backend commands.
- Tauri v2 for cross-platform desktop packaging and secure webview integration.
- Vanilla TypeScript, HTML, and CSS for the MVP frontend.
- A CLI bridge layer that starts with a fake command and later calls the existing Smara CLI.
- Local-first configuration and logs.

## Rationale

Tauri is preferred for the MVP because:

- It is much lighter than Electron.
- It gives a Rust-native backend layer.
- It supports Linux, Windows, and macOS.
- It can reuse a webview UI without shipping a full browser runtime.
- It allows gradual migration: Smara CLI can remain the source of truth while desktop-specific features are added in Rust.

## Alternatives considered

### Electron

Rejected for MVP because it usually produces larger bundles and higher memory usage.

### Wails

Rejected for this Rust-focused desktop direction because it is Go-based and overlaps with the previous Smara desktop experiment.

### Native Rust UI with egui/iced

Deferred. It may be useful later, but Tauri gives faster UI iteration and easier cross-platform packaging for the first MVP.

## Consequences

Positive:

- Lightweight desktop shell.
- Clear Rust backend boundary.
- Faster MVP delivery.
- Minimal disruption to existing Smara CLI.

Trade-offs:

- Webview behavior may differ between platforms.
- Tauri system dependencies must be handled in developer documentation.
- CLI bridge needs careful process handling and security review.

## Rollback criteria

Revisit this decision if:

- Tauri packaging becomes too fragile for target platforms.
- Webview limitations block core UX.
- Native UI becomes more important than webview-based UI.
- Bundle size or memory usage does not meet the lightweight target.
