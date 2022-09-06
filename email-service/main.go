package main

import (
	"log"
	"net/http"
	"os"
	"salazar/cheers/email/subscriptions"
)

func main() {
	http.HandleFunc("/", subscriptions.PaymentSub)

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
