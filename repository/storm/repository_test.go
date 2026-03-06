package storm

import (
	"path/filepath"
	"testing"
	"time"

	stormdb "github.com/asdine/storm/v3"

	"github.com/ajaxray/geek-life/model"
)

func openTestDB(t *testing.T) *stormdb.DB {
	t.Helper()

	db, err := stormdb.Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("open test db: %v", err)
	}

	t.Cleanup(func() {
		if closeErr := db.Close(); closeErr != nil {
			t.Fatalf("close test db: %v", closeErr)
		}
	})

	return db
}

func createProject(t *testing.T, repo interface {
	Create(title, uuid string) (model.Project, error)
}, title, uuid string) model.Project {
	t.Helper()

	project, err := repo.Create(title, uuid)
	if err != nil {
		t.Fatalf("create project: %v", err)
	}

	return project
}

func createTask(t *testing.T, repo interface {
	Create(project model.Project, title, details, uuid string, dueDate int64) (model.Task, error)
}, project model.Project, title, details, uuid string, dueDate int64) model.Task {
	t.Helper()

	task, err := repo.Create(project, title, details, uuid, dueDate)
	if err != nil {
		t.Fatalf("create task: %v", err)
	}

	return task
}

func mustTime(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
