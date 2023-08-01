package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {

		http.Error(w, "Methof not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(w, "POST Request is SuccessFull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name :%s\n", name)
	fmt.Fprintf(w, "Address :%s\n", address)
}

func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Starting server at Port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
