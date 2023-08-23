package main

import (
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/email-service/internal/app/events"
	"log"
	"net/http"
	"os"
)

func init() {
	go profiler.InitProfiling(os.Getenv("K_SERVICE"), "1.0.0")
}

func main() {
	http.HandleFunc("/", events.PaymentSub)
	http.HandleFunc("/user-sub", events.UserSub)

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
