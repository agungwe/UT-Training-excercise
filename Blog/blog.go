package blog

import (
	"fmt"
)

type BlogInterface interface {
	GetBlog()
	SetBlog(string, string, string)
}

type Blog struct {
	Title       string
	Descritpion string
	Author      string
}

func New(title string, desc string, author string) BlogInterface {
	return Blog{title, desc, author}
}

func (b Blog) GetBlog() {
	fmt.Println(b)
}

func (b Blog) SetBlog(title string, desc string, author string) {

}
