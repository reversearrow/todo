package tasks

import (
	"fmt"
	"testing"
)

func TestNewTask(t *testing.T) {
	todo, err := newTask("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(todo.createdAt.Clock())
}
