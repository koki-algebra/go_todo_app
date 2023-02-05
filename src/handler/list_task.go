package handler

import (
	"net/http"

	"github.com/koki-algebra/go_todo_app/entitiy"
	"github.com/koki-algebra/go_todo_app/store"
)

// ListTask
type ListTask struct {
	Store *store.TaskStore
}

type task struct {
	ID     entitiy.TaskID     `json:"id"`
	Title  string             `json:"title"`
	Status entitiy.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks := lt.Store.All()
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
