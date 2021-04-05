package registration

import "github.com/reversearrow/todo/app/user"

type Registration struct {
	Name     string
	Email    string
	Password string
	Verified bool
	UserData user.UserData
}

func NewRegistration(name, email, password string) {
	r := new(Registration)
	r.name = name
	r.email = email
	r.password = password
	r.verified = true
}
