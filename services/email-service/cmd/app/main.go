package main

import (
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/email-service/internal/app"
	"log"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("DISABLE_PROFILER") == "" {
		log.Println("Profiling enabled.")
		go profiler.InitProfiling("email-service", "1.0.0")
	} else {
		log.Println("Profiling disabled.")
	}

	http.HandleFunc("/", app.PaymentSub)

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
