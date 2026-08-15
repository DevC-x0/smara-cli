.PHONY: dev build run release-binaries test clean

# Default: Run Go MCP Server in dev mode
dev:
	go run ./cmd/server

# Build CLI / Server binary
build:
	@mkdir -p bin
	go build -o bin/smara ./cmd/server

# Run the compiled binary
run: build
	./bin/smara

# Cross-compile release binaries for Windows, macOS, Linux
release-binaries:
	@mkdir -p dist/bin dist/release
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o dist/bin/smara-linux-amd64 ./cmd/server
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o dist/bin/smara-linux-arm64 ./cmd/server
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o dist/bin/smara-darwin-amd64 ./cmd/server
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o dist/bin/smara-darwin-arm64 ./cmd/server
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o dist/bin/smara-windows-amd64.exe ./cmd/server
	@cd dist/bin && \
		tar -czvf ../release/smara-linux-amd64.tar.gz smara-linux-amd64 && \
		tar -czvf ../release/smara-linux-arm64.tar.gz smara-linux-arm64 && \
		tar -czvf ../release/smara-darwin-amd64.tar.gz smara-darwin-amd64 && \
		tar -czvf ../release/smara-darwin-arm64.tar.gz smara-darwin-arm64 && \
		zip ../release/smara-windows-amd64.zip smara-windows-amd64.exe && \
		cd ../release && sha256sum * > checksums.txt
	@echo "Release archives ready in dist/release/"

test:
	go test ./...

clean:
	rm -rf bin/ dist/
