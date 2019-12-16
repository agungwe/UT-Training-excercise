package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func main() {
	//"success": true,
	// String contains two JSON rows.
	text := "[{\"Id\": 1, \"Name\": \"User 1\"}, {\"Id\": 2, \"Name\": \"User 2\"}, {\"Id\": 3, \"Name\": \"User 3\"}]"
	// Get byte slice from string.

	bytes := []byte(text)

	// Unmarshal string into structs.
	var user []User
	json.Unmarshal(bytes, &user)

	// Loop over structs and display them.
	for l := range user {
		fmt.Printf("id_%v#Name_%v", user[l].Id, user[l].Name)
		fmt.Println()
	}
}
