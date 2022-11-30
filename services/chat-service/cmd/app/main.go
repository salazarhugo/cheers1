package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	websocket "github.com/gorilla/websocket"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	profiler "github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/chat-service/cache"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/app"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
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

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/ws", serveWs)

	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := app.NewServer()

	grpc_health_v1.RegisterHealthServer(grpcS, server)
	chat.RegisterChatServiceServer(grpcS, server)

	mixedHandler := newHTTPandGRPCMux(httpMux, grpcS)
	http2Server := &http2.Server{}
	http1Server := &http.Server{Handler: h2c.NewHandler(mixedHandler, http2Server)}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	err = http1Server.Serve(listener)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		panic(err)
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

	go WriteMessages(conn)
	go ReadMessages(conn)
}

func WriteMessages(conn *websocket.Conn) {
	cache := cache.NewCache(
		time.Duration(time.Duration.Hours(1)),
		redis.NewClient(&redis.Options{
			Addr:     "redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
			Password: "mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
			DB:       0,
		}),
	)
	repo := repository.NewChatRepository(cache, nil)
	rooms, err := repo.ListRoom("55FEvHawinQCa9jgH7ZdWESR3ri2")
	bytes, err := json.Marshal(rooms)

	err = conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		log.Println(err)
	}
}

func ReadMessages(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("failed to read message %v", err)
			conn.Close()
			return
		}
		log.Println(string(msg))
	}
}

func newHTTPandGRPCMux(httpHand http.Handler, grpcHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("content-type"), "application/grpc") {
			log.Println("GRPC")
			grpcHandler.ServeHTTP(w, r)
			return
		}
		log.Println("HTTP")
		httpHand.ServeHTTP(w, r)
	})
}
