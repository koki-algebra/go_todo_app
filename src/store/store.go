package store

import (
	"context"
	"errors"

	"github.com/koki-algebra/go_todo_app/entitiy"
)

var (
	Tasks = &TaskStore{Tasks: map[entitiy.TaskID]*entitiy.Task{}}

	ErrNotFound = errors.New("not found")
)

// TaskStore
type TaskStore struct {
	LastID entitiy.TaskID
	Tasks  map[entitiy.TaskID]*entitiy.Task
}

func (ts *TaskStore) Add(t *entitiy.Task) (entitiy.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) All() entitiy.Tasks {
	tasks := make([]*entitiy.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}
	return tasks
}

func (r *Respository) ListTasks(ctx context.Context, db Queryer) (entitiy.Tasks, error) {
	tasks := entitiy.Tasks{}
	sql := `SELECT id, title, status, created, modified FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Respository) AddTask(ctx context.Context, db Execer, t *entitiy.Task) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task (title, status, created, modified) VALUES (?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, t.Title, t.Status, t.Created, t.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entitiy.TaskID(id)

	return nil
}
