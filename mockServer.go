package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var requestCount int
var mu sync.Mutex
var resetTime time.Time

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Check if a minute has passed since the last reset
	if time.Since(resetTime) > time.Minute {
		requestCount = 0
		resetTime = time.Now()
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	requestCount++

	if requestCount > 100 {
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		return
	}

	// Handle the request (e.g., log it, store it, etc.)

	// Send a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
}

func main() {
	resetTime = time.Now()

	http.HandleFunc("/", handler)

	fmt.Println("Mock server listening on port 4566...")
	log.Fatal(http.ListenAndServe(":4566", nil))
}
