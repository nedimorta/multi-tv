package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// Channel represents a YouTube channel
type Channel struct {
	Name string `json:"name"`
	ID   string `json:"id"` // YouTube Channel ID
}

// AppState holds the application state
type AppState struct {
	sync.Mutex
	Channels []Channel
}

var appState AppState

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/channels", handleChannels)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handleIndex serves the main page
func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/public/index.html")
}

// handleChannels handles the addition and retrieval of YouTube channels
func handleChannels(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(appState.Channels)
	case http.MethodPost:
		var newChannel Channel
		if err := json.NewDecoder(r.Body).Decode(&newChannel); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		appState.Lock()
		appState.Channels = append(appState.Channels, newChannel)
		appState.Unlock()
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
