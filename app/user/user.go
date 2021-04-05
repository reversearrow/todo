package user

import (
	"github.com/google/uuid"
	"github.com/reversearrow/todo/app/todo"
)

type User struct {
	Username string
	UserData UserData
}

type UserData struct {
	Id       uuid.UUID
	TaskList todo.Todo
}

func NewUser() *UserData {
	u := new(UserData)
	u.Id = uuid.New()
	return u
}
