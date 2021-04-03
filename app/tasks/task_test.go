package tasks

import (
	"fmt"
	"testing"
)

func TestNewTask(t *testing.T) {
	todo, err := NewTask("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(todo.CreatedAt.Clock())
}
