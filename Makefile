.PHONY: dev dev-desktop dev-server build build-server build-desktop test clean

# Default: Run desktop dev app
dev: dev-desktop

# Run Smara Desktop dev mode (Tauri + Rust + Vite)
dev-desktop:
	$(MAKE) -C smara-desktop-rust dev

# Run Go MCP Server in dev mode
dev-server:
	go run ./cmd/server

# Build binaries
build: build-server build-desktop

build-server:
	go build -o bin/mcp-shared-memory-go ./cmd/server

build-desktop:
	$(MAKE) -C smara-desktop-rust build

test:
	go test ./...
	$(MAKE) -C smara-desktop-rust test

clean:
	rm -rf bin/
	$(MAKE) -C smara-desktop-rust clean
