package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	indexHandlerF := http.HandlerFunc(indexHandler)
	mux.Handle("GET /", withLogging(indexHandlerF))

	log.Print("Starting simple app on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatalf("Server error: %v", err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	msg := welcome("Adventurer")
	fmt.Fprint(w, msg)
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"Method: %s, Path: %s, Duration: %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
