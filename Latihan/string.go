package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonStr := `[{
		"Id": 1,
		"Name": "User 1"
	}, {
		"Id": 2,
		"Name": "User 2"
	}, {
		"Id": 3,
		"Name": "User 3"
	}]`

	type Person struct {
		Id   string
		Name string
	}
	var people []Person

	var personMap []map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &personMap)

	if err != nil {
		panic(err)
	}

	for _, personData := range personMap {

		// convert map to array of Person struct
		var p Person
		p.Id = fmt.Sprintf("%s", personData["Id"])
		p.Name = fmt.Sprintf("%s", personData["Name"])
		people = append(people, p)

	}
	fmt.Println(people)

	for _, value := range personMap {
		fmt.Println("Id:", value["Id"], "Name:", value["Name"])
	}
}
