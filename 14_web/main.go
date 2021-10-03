package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(
		writer,
		"<h1>Hello World<h2>",
	)
}

func about(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(
		writer,
		"<h1>About<h2>",
	)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting")
	_ = http.ListenAndServe(":3000", nil)

}
