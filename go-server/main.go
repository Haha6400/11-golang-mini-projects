package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfully \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method isnt supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	//Create a file server to serve static files form "/static"
	fileServer := http.FileServer(http.Dir("./static"))
	//Set a default procesing for the root ("/").
	http.Handle("/", fileServer)
	//Create a handle for the path "/form" and spacifies the formHandler function to handle requests
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at 8080\n")
	//Start server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
