package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request received")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}