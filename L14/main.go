package main

import "fmt"

func main() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}
	b := [][]byte{}

	for _, item := range items {
		a = append(a, item[:])
	}

	fmt.Println(items)
	fmt.Println(a)

	/*
		[[1 2] [3 4] [5 6]]
		[[1 2] [3 4] [5 6]]
		[[5 6] [5 6] [5 6]] bug is fixed in Go 1.22 :)
	*/

	for _, item := range items {
		i := make([]byte, len(item))
		copy(i, item[:]) // make unique slice for iteration
		b = append(b, item[:])
	}

	fmt.Println(items)
	fmt.Println(b)
}
