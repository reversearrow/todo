package registration

type Users struct {
	Users []*User
}

type User struct {
	Name     string
	Password string
}

func NewRegistration(name, password string) *User {
	r := new(User)
	r.Name = name
	r.Password = password
	return r
}

func (u *Users) AddUser(user *User) {
	u.Users = append(u.Users, user)
}
