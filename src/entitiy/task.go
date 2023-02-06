package entitiy

import "time"

// TaskID
type TaskID int64

// TaskStatus
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

// Task
type Task struct {
	ID       TaskID     `json:"id" db:"id"`
	Title    string     `json:"title" db:"title"`
	Status   TaskStatus `json:"status" db:"status"`
	Created  time.Time  `json:"created" db:"created"`
	Modified time.Time  `json:"modified" db:"modified"`
}

// Tasks
type Tasks []*Task
