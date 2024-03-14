package main

import "fmt"

// diubah menjadi kapital
type User struct {
	Email    string
	Password string
}

// diubah menjadi kapital juga ditambah dengan s dibelakang, penanda tidak hanya 1
type UserRepos struct {
	DB []User
}

func (r UserRepos) Register(u User) {
	if u.Email == "" || u.Password == "" {
		fmt.Println("register failed")
	}

	r.DB = append(r.DB, u)
}

func (r UserRepos) Login(u User) string {
	if u.Email == "" || u.Password == "" {
		fmt.Println("login failed")
	}

	// dari us menjadi user agar lebih jelas
	for _, user := range r.DB {
		if user.Email == u.Email && user.Password == u.Password {
			return "auth token"
		}
	}

	return ""
}