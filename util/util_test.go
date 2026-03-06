package util

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestConnectStormWithExplicitPath(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "nested", "tasks.db")

	db, err := ConnectStorm(dbPath)
	if err != nil {
		t.Fatalf("ConnectStorm returned error: %v", err)
	}

	t.Cleanup(func() {
		if closeErr := db.Close(); closeErr != nil {
			t.Fatalf("close db: %v", closeErr)
		}
	})

	if _, err := os.Stat(dbPath); err != nil {
		t.Fatalf("expected database file to exist at %q: %v", dbPath, err)
	}
}

func TestConnectStormRejectsDirectoryPath(t *testing.T) {
	dir := t.TempDir()

	db, err := ConnectStorm(dir)
	if err == nil {
		if db != nil {
			_ = db.Close()
		}
		t.Fatal("expected an error when database path is a directory")
	}
}

func TestUnixToTime(t *testing.T) {
	got := UnixToTime("1709251200")
	want := time.Unix(1709251200, 0)
	if !got.Equal(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestUnixToTimeFallsBackToCurrentTime(t *testing.T) {
	before := time.Now().Add(-2 * time.Second)
	got := UnixToTime("not-a-timestamp")
	after := time.Now().Add(2 * time.Second)

	if got.Before(before) || got.After(after) {
		t.Fatalf("expected fallback time between %v and %v, got %v", before, after, got)
	}
}
