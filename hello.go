package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json: "name"`
	Age  string `json: "age"`
}

func main() {
	var JsonString = `{"name":"Agung","age":"37"}`
	var JsonData = []byte(JsonString)

	var data User

	var err = json.Unmarshal(JsonData, &data)

	if err != nil {
		fmt.Println("Error Unmarshalling Json" + err.Error())
		return
	}

	fmt.Println(data)

	fmt.Println("My Name is " + data.Name)
	fmt.Println("My Age is " + data.Age)
}
