package main

import (
	"encoding/json"
	"net/http"
)

// Channel represents a YouTube channel
type Channel struct {
	Name string `json:"name"`
	ID   string `json:"id"` // YouTube Channel ID
}

// Channels stores the collection of Channel
var Channels []Channel

// handleIndex serves the main page
func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/build/index.html")
}

// handleChannels handles the addition and retrieval of YouTube channels
func handleChannels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow cross-origin requests

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(Channels)
	case http.MethodPost:
		var newChannel Channel
		if err := json.NewDecoder(r.Body).Decode(&newChannel); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		Channels = append(Channels, newChannel)
		json.NewEncoder(w).Encode(newChannel)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// handleChannelCount updates the number of channels displayed
func handleChannelCount(w http.ResponseWriter, r *http.Request) {
	// Placeholder for the implementation
}
