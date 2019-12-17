package main

import (
	"go/product"
	"go/profile"
)

func main() {
	var name = "Agung"
	profile := profile.User{}

	profile.SetProfile(name, "Kamu", "Dimana?")
	profile.GetProfile()

	product := product.New("Aqua", 100)
	product.GetProduct()
}
