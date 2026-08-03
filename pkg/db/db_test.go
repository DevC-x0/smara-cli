package db

import (
	"testing"
)

func TestDBOperations(t *testing.T) {
	database, err := InitDB()
	if err != nil {
		t.Fatalf("InitDB failed: %v", err)
	}
	defer database.Close()

	testPath := "/tmp/test-project"

	// Test SaveSession & GetLatestSession
	id, err := database.SaveSession(testPath, "Test summary", "Step 1", "main.go")
	if err != nil || id == 0 {
		t.Fatalf("SaveSession failed: %v", err)
	}

	sess, err := database.GetLatestSession(testPath)
	if err != nil || sess.Summary != "Test summary" {
		t.Fatalf("GetLatestSession failed: %v", err)
	}

	// Test StoreKnowledge & QueryKnowledge
	kId, err := database.StoreKnowledge(testPath, "AuthModule", "Architecture", "Uses JWT")
	if err != nil || kId == 0 {
		t.Fatalf("StoreKnowledge failed: %v", err)
	}

	results, err := database.QueryKnowledge(testPath, "JWT")
	if err != nil || len(results) == 0 {
		t.Fatalf("QueryKnowledge failed: %v", err)
	}
}
