package main

import "fmt"

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
}

func (u User) String() string {
	return fmt.Sprintf("{"+
		"\"name\": %s,"+
		"\"age\": %d,"+
		"\"email\": %s}", u.Name, u.Age, u.Email)
}
