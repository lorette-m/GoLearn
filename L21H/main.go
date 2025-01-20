package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func main() {
	db := database{"silly cat": 10, "smart cat": 20, "funny cat": 15}
	// add some routes
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/read", db.fetch)
	http.HandleFunc("/delete", db.remove)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// add the handlers

func (db database) list(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Pet shop!!\n")
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	cost := req.URL.Query().Get("cost")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}

	c, err := strconv.ParseFloat(cost, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", cost)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}

	db[item] = dollars(c)

	fmt.Fprintf(w, "added %s with price %s\n", item, db[item])
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	cost := req.URL.Query().Get("cost")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	c, err := strconv.ParseFloat(cost, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", cost)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}

	db[item] = dollars(c)

	fmt.Fprintf(w, "new price %s for item %s\n", db[item], item)
}

func (db database) fetch(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	fmt.Fprintf(w, "item %s has price %s\n", item, db[item])
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "item %s is removed\n", item)
}
