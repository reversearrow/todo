package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type users struct{
	users []*User
}

var u users

func(u *users) addUser(user *User){
	for _, v := range u.users{
		if v.Email == user.Email{
			return
		}
	}
	u.users = append(u.users, user)
}

type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
}

func registration(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		for _, v := range u.users{
			fmt.Println("users")
			fmt.Println(v.Email)
		}
		w.WriteHeader(http.StatusOK)
	case "POST":
		b := make([]byte, r.ContentLength)
		r.Body.Read(b)
		var user User
		json.Unmarshal(b, &user)
		u.addUser(&user)
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", registration)
	server := http.Server{
		Handler: mux,
	}
	server.ListenAndServe()
}
