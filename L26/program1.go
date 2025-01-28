package main

import "fmt"

func program1() {
	//ch := make(chan int, 1)
	ch := make(chan int)

	ch <- 1

	b, ok := <-ch
	fmt.Println(b, ok) // 1 true

	close(ch)

	c, ok := <-ch
	fmt.Println(c, ok) // 0 false
}
