package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	} `json:"name"`
}

func main() {
	var wg sync.WaitGroup
	userCh := make(chan User)

	wg.Add(1)
	go func() {
		defer wg.Done()
		getUsers := GetRequest()

		for _, user := range getUsers {
			userCh <- user
		}

		close(userCh)
	}()

	fmt.Println("Users:")
	for user := range userCh {
		fmt.Println("===")
		fmt.Printf("ID: %v\nUsername: %v\nEmail: %v\nFirst Name: %v\nLast Name: %v\n", user.ID, user.Username, user.Email, user.Name.FirstName, user.Name.LastName)
		fmt.Println("===")
	}

	wg.Wait()
}

func GetRequest() []User {
	var users []User
	c := &http.Client{Timeout: 10 * time.Second}

	r, err := http.NewRequest("GET", "https://fakestoreapi.com/users", nil)
	if err != nil {
		log.Println(err)
	}

	res, err := c.Do(r)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&users)

	return users
}
