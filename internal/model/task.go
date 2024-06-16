package model

import (
	"context"
	"time"

	"github.com/symonk/toodoo/internal/db"
)

const (
	fetchAllTasks = "SELECT * FROM task;"
	fetchTask     = "SELECT * FROM task WHERE id = ? LIMIT 1"
)

// TaskModel is the core encapsulation of a task
type TaskModel struct {
	Id          int       `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Recurring   bool      `json:"recurring"`
	Schedule    time.Time `json:"schedule"`
}

func NewTask() *TaskModel {
	return &TaskModel{
		Name:        "foo",
		Description: "bar",
		Recurring:   true,
		Schedule:    time.Now().Add(time.Minute),
	}
}

// RetrieveTasks fetches all tasks from the database.
func (t TaskModel) RetrieveTasks(ctx context.Context) ([]TaskModel, error) {
	tasks := make([]TaskModel, 0)
	client := db.GetDB()
	if err := client.SelectContext(ctx, &tasks, db.RebindQuery(client.DB, fetchAllTasks)); err != nil {
		return tasks, err
	}
	return tasks, nil
}

// RetrieveTaskById returns the task with the given ID.
func (t TaskModel) RetrieveTaskByID(ctx context.Context, id int) (TaskModel, error) {
	var task TaskModel
	client := db.GetDB()
	if err := client.GetContext(ctx, &task, db.RebindQuery(client.DB, fetchTask), id); err != nil {
		return task, err
	}
	return task, nil
}
