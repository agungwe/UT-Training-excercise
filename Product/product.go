package product

import "fmt"

type Product struct {
	ProductName string
	Qty  int32
}

func New(name string, qty int32) Product {
	p := Product{
		ProductName: name,
		Qty: qty,
	}

	return p
}

func (p Product) GetProduct() {
	fmt.Println(p)
}