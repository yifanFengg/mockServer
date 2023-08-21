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
var successfulRequest int
var errorCount int

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Check if a minute has passed since the last reset
	//if time.Since(resetTime) > 10*time.Second {
	//	requestCount = 0
	//	resetTime = time.Now()
	//}

	requestCount++

	//if requestCount%100 == 0 {
	//	errorCount = 0
	//}
	//if r.Method != http.MethodPost {
	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//	return
	//}

	//if r.URL.Path != "/" {
	//	http.Error(w, "Not found", http.StatusNotFound)
	//	return
	//}

	if requestCount%100 < 80 {
		http.Error(w, "too many requests", http.StatusTooManyRequests)
		log.Println("Too many requests,wait for a while")
		return
	}

	// Handle the request (e.g., log it, store it, etc.)

	// Log the request
	//log.Println("Received POST request:", r.URL.Path)
	//log.Println("request count: %v", requestCount)

	// Send a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success"}`))
	successfulRequest++
	log.Println("successfully upload count: %v", successfulRequest)
	log.Println()
}

func main() {
	// Open a log file
	successfulRequest = 0
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Set log output to the file
	log.SetOutput(logFile)

	//resetTime = time.Now()

	http.HandleFunc("/", handler)

	log.Println("Mock server listening on port 4567...")
	log.Fatal(http.ListenAndServe(":4567", nil))
}
