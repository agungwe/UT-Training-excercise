package main

import (
	"fmt"
)

type User struct {
	Name    string
	Address string
	Age     string
}

func main() {

	var user = User{}
	user.Name = "AgungW"
	user.Address = "Jakarta"

	var user1 = User{"Agung W", "Jakarta", "37"}

	var user2 = User{Name: "Agung W", Address: "Purwokerto"}

	fmt.Println(user)
	fmt.Println(user1)
	fmt.Println(user2)
}
