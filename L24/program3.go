package main

import (
	"log"
	"time"
)

func program3() {
	const tickRate = 2 * time.Second

	stopper := time.After(5 * tickRate)
	ticker := time.NewTicker(tickRate).C
	log.Println("start")
loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			break loop
		}
	}
	
	log.Println("finish")
}
