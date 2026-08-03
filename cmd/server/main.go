package main

import (
	"log"

	"mcp-bridge-memory/pkg/db"
	mcpServer "mcp-bridge-memory/pkg/mcp"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to init database: %v", err)
	}
	defer database.Close()

	s := server.NewMCPServer("mcp-shared-memory-go", "1.0.0")
	mcpServer.RegisterTools(s, database)

	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
