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

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() error %v", err)
		return
	}

	fmt.Fprintf(writer, "Post request successful\n")

	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(writer, "name = %s\n", name)
	fmt.Fprintf(writer, "address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Server is running on port 8000. http://localhost:8000/\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
