package main

import (
	"github.com/salazarhugo/cheers1/services/notification-service/internal/app"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", app.ChatEventPubSub)
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	// Start HTTP server.
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
