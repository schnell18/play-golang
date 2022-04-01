package main

import (
	"fmt"
	"net/http"
	"os"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request succeed\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Listening on port 8000 ...\n")
	fmt.Printf("Try http://127.0.0.1:8000/\n")
	fmt.Printf("Try http://127.0.0.1:8000/form.html\n")
	fmt.Printf("Try http://127.0.0.1:8000/hello\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Fail to ListenAndServe() due to: %v", err)
		os.Exit(1)
	}
}
