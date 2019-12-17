package main

import (
	//"go/profile"
	"go/product"
)

func main() {
	//var name = "Agung"
	//profile := profile.User{}

	//profile.GetProfile()
	//profile.SetProfile(name, "Kamu", "Dimana?")

	product := product.New("Aqua", 100)
	product.GetProduct()
}
