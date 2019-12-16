package main

import "fmt"

func main() {
	var jenis = []string{"foo1", "bar1", "foo2", "bar2"}

	for _, loop := range jenis {
		fmt.Println(loop)
	}

}
