package main

import (
	"github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/email-service/internal/app/events"
	"log"
	"net/http"
	"os"
)

func main() {
	go profiler.InitProfiling("email-service", "1.0.0")

	http.HandleFunc("/", events.PaymentSub)
	http.HandleFunc("/auth-sub", events.AuthSub)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
		log.Printf("Defaulting to port %s", port)
	}

	// Start HTTP server.
	//log.Printf("Listening on port %s", port)
	//if err := http.ListenAndServe(":"+port, nil); err != nil {
	//	log.Fatal(err)
	//}

	err := events.SendEmail("hugobrock74@gmail.com", "Hugo", []*ticket.Ticket{}, 12000)
	if err != nil {
		log.Println(err)
		return
	}
}
