package storm

import (
	"testing"
	"time"
)

func TestTaskRepositoryCRUDAndQueries(t *testing.T) {
	db := openTestDB(t)
	projectRepo := NewProjectRepository(db)
	taskRepo := NewTaskRepository(db)
	project := createProject(t, projectRepo, "Inbox", "project-uuid")

	today := mustTime(2026, 3, 6)
	tomorrow := today.AddDate(0, 0, 1)

	unscheduled := createTask(t, taskRepo, project, "Unscheduled", "No due date", "task-1", 0)
	dueToday := createTask(t, taskRepo, project, "Today", "Due today", "task-2", today.Unix())
	dueTomorrow := createTask(t, taskRepo, project, "Tomorrow", "Due tomorrow", "task-3", tomorrow.Unix())

	all, err := taskRepo.GetAll()
	if err != nil {
		t.Fatalf("GetAll returned error: %v", err)
	}
	if len(all) != 3 {
		t.Fatalf("expected 3 tasks, got %d", len(all))
	}

	byID, err := taskRepo.GetByID(dueToday.ID)
	if err != nil {
		t.Fatalf("GetByID returned error: %v", err)
	}
	if byID.UUID != dueToday.UUID {
		t.Fatalf("expected UUID %q, got %q", dueToday.UUID, byID.UUID)
	}

	byUUID, err := taskRepo.GetByUUID(dueTomorrow.UUID)
	if err != nil {
		t.Fatalf("GetByUUID returned error: %v", err)
	}
	if byUUID.Title != dueTomorrow.Title {
		t.Fatalf("expected title %q, got %q", dueTomorrow.Title, byUUID.Title)
	}

	byProject, err := taskRepo.GetAllByProject(project)
	if err != nil {
		t.Fatalf("GetAllByProject returned error: %v", err)
	}
	if len(byProject) != 3 {
		t.Fatalf("expected 3 tasks for project, got %d", len(byProject))
	}

	unscheduledTasks, err := taskRepo.GetAllByDate(time.Time{})
	if err != nil {
		t.Fatalf("GetAllByDate zero-value query returned error: %v", err)
	}
	if len(unscheduledTasks) != 1 || unscheduledTasks[0].ID != unscheduled.ID {
		t.Fatalf("expected unscheduled task %d, got %#v", unscheduled.ID, unscheduledTasks)
	}

	todayTasks, err := taskRepo.GetAllByDate(today)
	if err != nil {
		t.Fatalf("GetAllByDate returned error: %v", err)
	}
	if len(todayTasks) != 1 || todayTasks[0].ID != dueToday.ID {
		t.Fatalf("expected due-today task %d, got %#v", dueToday.ID, todayTasks)
	}

	rangeTasks, err := taskRepo.GetAllByDateRange(today, tomorrow)
	if err != nil {
		t.Fatalf("GetAllByDateRange returned error: %v", err)
	}
	if len(rangeTasks) != 2 {
		t.Fatalf("expected 2 ranged tasks, got %d", len(rangeTasks))
	}

	if err := taskRepo.UpdateField(&dueToday, "Completed", true); err != nil {
		t.Fatalf("UpdateField returned error: %v", err)
	}

	updated, err := taskRepo.GetByID(dueToday.ID)
	if err != nil {
		t.Fatalf("GetByID after UpdateField returned error: %v", err)
	}
	if !updated.Completed {
		t.Fatal("expected task to be marked completed")
	}

	if err := taskRepo.Delete(&dueTomorrow); err != nil {
		t.Fatalf("Delete returned error: %v", err)
	}

	all, err = taskRepo.GetAll()
	if err != nil {
		t.Fatalf("GetAll after Delete returned error: %v", err)
	}
	if len(all) != 2 {
		t.Fatalf("expected 2 tasks after delete, got %d", len(all))
	}
}
