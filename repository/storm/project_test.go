package storm

import "testing"

func TestProjectRepositoryCRUD(t *testing.T) {
	db := openTestDB(t)
	repo := NewProjectRepository(db)

	created, err := repo.Create("Inbox", "project-uuid")
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}

	byID, err := repo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("GetByID returned error: %v", err)
	}
	if byID.Title != created.Title {
		t.Fatalf("expected title %q, got %q", created.Title, byID.Title)
	}

	byTitle, err := repo.GetByTitle(created.Title)
	if err != nil {
		t.Fatalf("GetByTitle returned error: %v", err)
	}
	if byTitle.UUID != created.UUID {
		t.Fatalf("expected UUID %q, got %q", created.UUID, byTitle.UUID)
	}

	byUUID, err := repo.GetByUUID(created.UUID)
	if err != nil {
		t.Fatalf("GetByUUID returned error: %v", err)
	}
	if byUUID.ID != created.ID {
		t.Fatalf("expected ID %d, got %d", created.ID, byUUID.ID)
	}

	if err := repo.UpdateField(&created, "Title", "Work"); err != nil {
		t.Fatalf("UpdateField returned error: %v", err)
	}

	updated, err := repo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("GetByID after UpdateField returned error: %v", err)
	}
	if updated.Title != "Work" {
		t.Fatalf("expected updated title %q, got %q", "Work", updated.Title)
	}

	all, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll returned error: %v", err)
	}
	if len(all) != 1 {
		t.Fatalf("expected 1 project, got %d", len(all))
	}

	if err := repo.Delete(&created); err != nil {
		t.Fatalf("Delete returned error: %v", err)
	}

	all, err = repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll after Delete returned error: %v", err)
	}
	if len(all) != 0 {
		t.Fatalf("expected 0 projects after delete, got %d", len(all))
	}
}
