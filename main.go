package main

import (
	"go/blog"
	"go/product"
	"go/profile"
	"go/store"
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

	branchStore := store.BranchStore{StoreName: "Cabang Toko Bakpia", OwnerBranch: "Agung Wicaksono"}
	store := store.Store{BranchStore: branchStore, StoreName: "Toko Bakpia", Owner: "Dzakiy"}

	store.GetBranchStore()
}
