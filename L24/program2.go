package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func program2() {
	stopper := time.After(3 * time.Second)
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
		"http://localhost:8080",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		select {
		case r := <-results:
			if r.err != nil {
				log.Printf(" %-20s %s\n", r.url, r.err)
			} else {
				log.Printf(" %-20s %s\n", r.url, r.latency)
			}

		case t := <-stopper:
			log.Fatalf("program2 timed out %s", t)
		}
	}
}
