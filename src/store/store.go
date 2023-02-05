package store

import (
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
