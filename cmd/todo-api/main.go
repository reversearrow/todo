package main

import (
	"fmt"

	"github.com/reversearrow/todo/app/todo"
)

func main() {
	var todo todo.Todo
	todo.Add("Hello1")
	todo.Add("Hello2")
	todo.Add("Hello3")
	for _, v := range todo.Tasks {
		fmt.Println(*v.Task)
	}
	id := todo.Tasks[0].Id
	todo.Update(id, "Hello Hello Hello")

	fmt.Println("After update")
	for _, v := range todo.Tasks {
		fmt.Println(*v.Task)
	}

	fmt.Println("Removing by ID")
	todo.Delete(id)

	for _, v := range todo.Tasks {
		fmt.Println(*v.Task)
	}
}
