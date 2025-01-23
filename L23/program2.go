package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//var nextID = make(chan int)

type nextCh chan int

func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(7) * time.Second)
	fmt.Fprintf(w, "<h1>You got %d<h1>", <-ch)

	//nextID++ // UNSAFE
}

func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func program2() {
	var nextID nextCh = make(chan int)

	go counter(nextID)
	http.HandleFunc("/", nextID.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
