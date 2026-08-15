# MCP

MCP atau Model Context Protocol memungkinkan Smara memakai tool eksternal sebagai bagian dari workflow agen.

## Auto-discovery

Smara dapat auto-load MCP server dari beberapa sumber:

- Konfigurasi Windsurf IDE.
- Konfigurasi OpenCode.
- Konfigurasi native Smara.
- Remote MCP via SSE/WebSocket.

## Kenapa MCP penting?

MCP membuat Smara bisa terhubung ke tool khusus seperti browser automation, database inspector, knowledge service, docs fetcher, atau integrasi internal perusahaan.

## Workflow umum

1. Tambahkan MCP server ke config.
2. Jalankan `smara start`.
3. Smara memuat daftar tool.
4. Agen dapat memakai tool sesuai izin dan policy.

## Safety

Untuk tool yang bisa menulis data, menjalankan command, atau mengakses server, gunakan mode `plan` agar Smara meminta approval sebelum eksekusi.
