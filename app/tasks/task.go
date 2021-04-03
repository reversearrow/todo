package tasks

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task string

func NewTaskString(t string) (*Task, error) {
	var tk Task
	if len(t) > 64 {
		return nil, errors.New("failed to create task string")
	}

	tk = Task(t)
	return &tk, nil
}

//TaskData
type TaskData struct {
	Id        uuid.UUID
	Task      *Task
	CreatedAt time.Time
}

func NewTask(t string) (*TaskData, error) {
	tk := new(TaskData)
	tk.Id = uuid.New()
	taskString, err := NewTaskString(t)
	if err != nil {
		return nil, errors.New("failed to create new task")
	}
	tk.Task = taskString
	tk.CreatedAt = time.Now()
	return tk, nil
}
