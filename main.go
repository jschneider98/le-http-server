package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/.well-known/", http.FileServer(http.Dir("/var/www")))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Steadfast"))
	})

	log.Println("HTTP server listening on :80")
	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
