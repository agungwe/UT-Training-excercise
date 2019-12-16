package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama   string `json:"name"`
	Alamat string `json:"alamat"`
	Umur   string `json:"umur"`
}

func main() {
	var dataJson = []User{{"adit", "jogja", "22"}, {"agungw", "pwt", "35"}}
	var jsonData, err = json.Marshal(dataJson)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
