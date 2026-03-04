package main

import "fmt"

// defining a struct type
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	// Address Address `json:"address"` -> nested struct ex. (user.Address.Street)
	Address `json:"address"` // embedding Address struct into User struct (user.Street)
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

func main() {
	// structs
	// initializing a struct
	user := User{
		Username: "kaloka8",
		Email:    "halobro@example.com",
		IsActive: true,
		Address:  Address{Street: "123 Main Street", City: "Purworejo", State: "Central Java"},
	}
	fmt.Println(user)
	fmt.Println(user.State)
}
