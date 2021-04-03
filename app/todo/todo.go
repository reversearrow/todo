package todo

import (
	"log"

	"github.com/google/uuid"
	"github.com/reversearrow/todo/app/tasks"
)

type Todo struct {
	Tasks []*tasks.TaskData
}

func (t *Todo) Add(task string) {
	tk, err := tasks.NewTask(task)
	if err != nil {
		log.Print("error adding task")
	}
	t.Tasks = append(t.Tasks, tk)
}

func (t *Todo) Update(id uuid.UUID, description string) {
	td, err := tasks.NewTaskString(description)
	if err != nil {
		log.Print("error updating task")
	}
	for _, v := range t.Tasks {
		if v.Id == id {
			v.Task = td
		}
	}
}

func (t *Todo) Delete(id uuid.UUID) {
	var index int
	for i, v := range t.Tasks {
		if v.Id == id {
			index = i
		}
	}
	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
}
