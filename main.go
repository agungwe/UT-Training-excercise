package main

import (
	"go/profile"
)

func main() {
	var name = "Agung"
	profile := profile.User{}

	profile.GetProfile()
	profile.SetProfile(name, "Kamu", "Dimana?")
}
