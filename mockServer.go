package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	if time.Since(resetTime) > 5*time.Second {
		requestCount = 0
		resetTime = time.Now()
	}

	//if r.Method != http.MethodPost {
	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//	return
	//}

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	requestCount++

	if requestCount > 100 {
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		log.Println("Too many requests,wait for a while")
		return
	}

	// Handle the request (e.g., log it, store it, etc.)

	// Log the request
	log.Println("Received POST request:", r.URL.Path)
	log.Println("request count: %v", requestCount)

	// Send a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
}

func main() {
	// Open a log file
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Set log output to the file
	log.SetOutput(logFile)

	resetTime = time.Now()

	http.HandleFunc("/", handler)

	log.Println("Mock server listening on port 4566...")
	log.Fatal(http.ListenAndServe(":4566", nil))
}
