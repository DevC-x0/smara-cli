package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"mcp-bridge-memory/pkg/db"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *server.MCPServer, database *db.DB) {
	// save_session_handover
	s.AddTool(
		mcp.NewTool("save_session_handover",
			mcp.WithDescription("Menyimpan checkpoint ringkasan sesi ke tabel sessions"),
			mcp.WithString("project_path", mcp.Required(), mcp.Description("Path repositori proyek")),
			mcp.WithString("summary", mcp.Required(), mcp.Description("Ringkasan status sesi")),
			mcp.WithString("next_steps", mcp.Description("Langkah berikutnya")),
			mcp.WithString("files_modified", mcp.Description("Daftar file yang diubah")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			path, err := req.RequireString("project_path")
			if err != nil {
				return mcp.NewToolResultError("project_path is required"), nil
			}
			summary, err := req.RequireString("summary")
			if err != nil {
				return mcp.NewToolResultError("summary is required"), nil
			}
			nextSteps := req.GetString("next_steps", "")
			filesModified := req.GetString("files_modified", "")

			id, err := database.SaveSession(path, summary, nextSteps, filesModified)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return mcp.NewToolResultText(fmt.Sprintf("Session handover saved with ID %d", id)), nil
		},
	)

	// get_session_handover
	s.AddTool(
		mcp.NewTool("get_session_handover",
			mcp.WithDescription("Mengambil ringkasan sesi paling terakhir dari tabel sessions"),
			mcp.WithString("project_path", mcp.Required(), mcp.Description("Path repositori proyek")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			path, err := req.RequireString("project_path")
			if err != nil {
				return mcp.NewToolResultError("project_path is required"), nil
			}
			sess, err := database.GetLatestSession(path)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("No session found or error: %v", err)), nil
			}
			out, _ := json.MarshalIndent(sess, "", "  ")
			return mcp.NewToolResultText(string(out)), nil
		},
	)

	// store_knowledge
	s.AddTool(
		mcp.NewTool("store_knowledge",
			mcp.WithDescription("Menyimpan entitas atau keputusan arsitektur ke knowledge_nodes"),
			mcp.WithString("project_path", mcp.Required(), mcp.Description("Path repositori proyek")),
			mcp.WithString("entity", mcp.Required(), mcp.Description("Nama entitas")),
			mcp.WithString("category", mcp.Required(), mcp.Description("Kategori")),
			mcp.WithString("fact", mcp.Required(), mcp.Description("Fakta/keputusan")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			path, err := req.RequireString("project_path")
			if err != nil {
				return mcp.NewToolResultError("project_path is required"), nil
			}
			entity, err := req.RequireString("entity")
			if err != nil {
				return mcp.NewToolResultError("entity is required"), nil
			}
			category, err := req.RequireString("category")
			if err != nil {
				return mcp.NewToolResultError("category is required"), nil
			}
			fact, err := req.RequireString("fact")
			if err != nil {
				return mcp.NewToolResultError("fact is required"), nil
			}

			id, err := database.StoreKnowledge(path, entity, category, fact)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return mcp.NewToolResultText(fmt.Sprintf("Knowledge stored with ID %d", id)), nil
		},
	)

	// query_knowledge
	s.AddTool(
		mcp.NewTool("query_knowledge",
			mcp.WithDescription("Mencari entitas/fakta terdaftar berdasarkan nama entitas/kategori"),
			mcp.WithString("project_path", mcp.Required(), mcp.Description("Path repositori proyek")),
			mcp.WithString("query", mcp.Required(), mcp.Description("Kata kunci pencarian")),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			path, err := req.RequireString("project_path")
			if err != nil {
				return mcp.NewToolResultError("project_path is required"), nil
			}
			query, err := req.RequireString("query")
			if err != nil {
				return mcp.NewToolResultError("query is required"), nil
			}

			nodes, err := database.QueryKnowledge(path, query)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			out, _ := json.MarshalIndent(nodes, "", "  ")
			return mcp.NewToolResultText(string(out)), nil
		},
	)
}
