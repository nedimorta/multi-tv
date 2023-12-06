package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./frontend/build"))) // Serve static files
	http.HandleFunc("/api/channels", handleChannels)
	http.HandleFunc("/api/channelcount", handleChannelCount) // New endpoint

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
