package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

type Session struct {
	ID            int64  `json:"id"`
	ProjectPath   string `json:"project_path"`
	Summary       string `json:"summary"`
	NextSteps     string `json:"next_steps,omitempty"`
	FilesModified string `json:"files_modified,omitempty"`
	CreatedAt     string `json:"created_at"`
}

type Knowledge struct {
	ID          int64  `json:"id"`
	ProjectPath string `json:"project_path"`
	Entity      string `json:"entity"`
	Category    string `json:"category"`
	Fact        string `json:"fact"`
	CreatedAt   string `json:"created_at"`
}

func InitDB() (*DB, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dir := filepath.Join(home, ".mcp-shared-memory")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", filepath.Join(dir, "db.sqlite")+"?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)")
	if err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_path TEXT NOT NULL,
		summary TEXT NOT NULL,
		next_steps TEXT,
		files_modified TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS knowledge_nodes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_path TEXT NOT NULL,
		entity TEXT NOT NULL,
		category TEXT NOT NULL,
		fact TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(schema); err != nil {
		db.Close()
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) SaveSession(projectPath, summary, nextSteps, filesModified string) (int64, error) {
	res, err := db.Exec(
		"INSERT INTO sessions (project_path, summary, next_steps, files_modified) VALUES (?, ?, ?, ?)",
		projectPath, summary, nextSteps, filesModified,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *DB) GetLatestSession(projectPath string) (*Session, error) {
	row := db.QueryRow(
		"SELECT id, project_path, summary, COALESCE(next_steps,''), COALESCE(files_modified,''), created_at FROM sessions WHERE project_path = ? ORDER BY id DESC LIMIT 1",
		projectPath,
	)
	var s Session
	if err := row.Scan(&s.ID, &s.ProjectPath, &s.Summary, &s.NextSteps, &s.FilesModified, &s.CreatedAt); err != nil {
		return nil, err
	}
	return &s, nil
}

func (db *DB) StoreKnowledge(projectPath, entity, category, fact string) (int64, error) {
	res, err := db.Exec(
		"INSERT INTO knowledge_nodes (project_path, entity, category, fact) VALUES (?, ?, ?, ?)",
		projectPath, entity, category, fact,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *DB) QueryKnowledge(projectPath, query string) ([]Knowledge, error) {
	rows, err := db.Query(
		"SELECT id, project_path, entity, category, fact, created_at FROM knowledge_nodes WHERE project_path = ? AND (entity LIKE ? OR category LIKE ? OR fact LIKE ?)",
		projectPath, "%"+query+"%", "%"+query+"%", "%"+query+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Knowledge
	for rows.Next() {
		var k Knowledge
		if err := rows.Scan(&k.ID, &k.ProjectPath, &k.Entity, &k.Category, &k.Fact, &k.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, k)
	}
	return result, rows.Err()
}
