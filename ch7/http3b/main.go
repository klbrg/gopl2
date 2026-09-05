// Http5 is the http4 server with Go 1.22 method-and-wildcard patterns.
package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.PathValue("item")
	price, ok := db[item]
	if !ok {
		http.Error(w, fmt.Sprintf("no such item: %q", item), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /list", db.list)
	mux.HandleFunc("GET /price/{item}", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
