package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 Page not found.", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(writer, "Hello!\n")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server is running on port 8000. http://localhost:8000/\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
