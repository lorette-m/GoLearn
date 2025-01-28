package main

import (
	"fmt"
	"time"
)

type T struct {
	i byte
	b bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	ch <- t

	t.b = true
}

func program2() {
	vs := make([]T, 5)
	ch := make(chan *T, 5)

	for i := range vs {
		go send(i, ch)
	}

	time.Sleep(1 * time.Second)

	// copy quickly!
	for i := range vs {
		vs[i] = *<-ch
	}

	// print later
	for _, v := range vs {
		fmt.Println(v)
	}
}
