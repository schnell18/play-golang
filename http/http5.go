package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{items: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	log.Println("Listen and serve on localhost:8000")
	// use DefaultServeMux to simply handler function registration
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	sync.Mutex
	items map[string]dollars
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.items {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db.items[item]
	if !ok {
		msg := fmt.Sprintf("no such item %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	oldPrice, ok := db.items[item]
	if !ok {
		msg := fmt.Sprintf("no such item %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	priceStr, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		msg := fmt.Sprintf("bad price %s due to: %q\n", priceStr, err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	newPrice := dollars(priceStr)
	db.Lock()
	db.items[item] = newPrice
	defer db.Unlock()

	fmt.Fprintf(w, "Updated %s price from %s to %s\n", item, oldPrice, newPrice)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db.items[item]
	if ok {
		msg := fmt.Sprintf("Item %s alreay exists\n", item)
		http.Error(w, msg, http.StatusNotAcceptable)
		return
	}
	priceStr, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		msg := fmt.Sprintf("bad price %s due to: %q\n", priceStr, err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	newPrice := dollars(priceStr)
	db.Lock()
	db.items[item] = newPrice
	defer db.Unlock()

	fmt.Fprintf(w, "Created item %s with price %s\n", item, newPrice)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db.items[item]
	if !ok {
		msg := fmt.Sprintf("no such item %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	db.Lock()
	delete(db.items, item)
	defer db.Unlock()

	fmt.Fprintf(w, "Deleted item %s\n", item)
}
