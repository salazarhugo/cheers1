package main

import (
	websocket "github.com/gorilla/websocket"
	profiler "github.com/salazarhugo/cheers1/libs/profiler"
	"log"
	"net/http"
	"os"
)

func init() {
	go profiler.InitProfiling("chat-service", "1.0.0")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}

	http.HandleFunc("/ws", serveWs)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

}
