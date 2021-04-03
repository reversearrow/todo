package tasks

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type task string

func newTaskString(t string) (*task, error) {
	var tk task
	if len(t) > 64 {
		return nil, errors.New("failed to create task string")
	}

	tk = task(t)
	return &tk, nil
}

//TaskData
type TaskData struct {
	id        uuid.UUID
	task      *task
	createdAt time.Time
}

func newTask(t string) (*TaskData, error) {
	tk := new(TaskData)
	tk.id = uuid.New()
	taskString, err := newTaskString(t)
	if err != nil {
		return nil, errors.New("failed to create new task")
	}
	tk.task = taskString
	tk.createdAt = time.Now()
	return tk, nil
}
