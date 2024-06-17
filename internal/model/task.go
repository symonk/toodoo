package model

import (
	"context"
	"time"

	"github.com/symonk/toodoo/internal/db"
)

const (
	fetchAllTasks = "SELECT * FROM task;"
	fetchTask     = "SELECT * FROM task WHERE id = ? LIMIT 1"
	createTask    = "INSERT INTO task (name, description, recurring, schedule) VALUES(?,?,?,?) RETURNING id"
)

// TaskModel is the core encapsulation of a task
type TaskModel struct {
	Id          int       `json:"id" binding:"-"`
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

// Create creates a new task instance in the database and returns the server
// side task that was just created.
func (t TaskModel) Create(ctx context.Context, task TaskModel) (TaskModel, error) {
	client := db.GetDB()
	var createdID int
	row := client.QueryRowContext(ctx, db.RebindQuery(client.DB, createTask), task.Name, task.Description, task.Recurring, task.Schedule)
	var newTask TaskModel

	if err := row.Scan(&createdID); err != nil {
		return newTask, err
	}

	if err := client.GetContext(ctx, &newTask, db.RebindQuery(client.DB, fetchTask), createdID); err != nil {
		return newTask, err
	}
	return newTask, nil

}
