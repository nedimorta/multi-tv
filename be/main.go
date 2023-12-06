// main.go
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	// Add more routes here

	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Handle your root route
}
