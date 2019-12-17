package main

import (
	"go/blog"
	"go/product"
	"go/profile"
)

func main() {
	var name = "Agung"
	profile := profile.User{}
	profile.GetProfile()
	profile.SetProfile(name, "Kamu", "Dimana?")

	product := product.New("Aqua", 100)
	product.GetProduct()

	blog := blog.New("Tutorial Golang", "OOP in", "Agung W")
	blog.GetBlog()

}
