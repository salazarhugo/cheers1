package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	mauth "github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/app"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/ws"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

var redisClient *redis.Client
var clientOnce sync.Once

// init runs during package initialization.
// this will only run during an instance's cold start.
func init() {
	go profiler.InitProfiling("chat-service", "1.0.0")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ENDPOINT"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       0,
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("Defaulting to port %s", port)
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/chat", serveWs)
	httpMux.HandleFunc("/user-sub", app.UserSub)

	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(mauth.UnaryInterceptor),
	)

	server := app.NewServer(redisClient)

	grpc_health_v1.RegisterHealthServer(grpcS, server)
	chat.RegisterChatServiceServer(grpcS, server)

	mixedHandler := utils.NewHttpAndGrpcMux(httpMux, grpcS)
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
	// Parse query parameters
	query := r.URL.Query()
	authToken := query.Get("token")

	// Validate JWT token
	token, err := ValidateJwt(authToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := token.UID

	h := http.Header{}
	h.Set("Sec-WebSocket-Protocol", userID)

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, h)
	if err != nil {
		log.Println(err)
	}
	connection := new(ws.Connection)
	connection.Socket = conn

	log.Println("WS: " + userID)

	go ws.SubscribePubSub(
		connection,
		redisClient,
		userID,
	)
	go ws.ReadMessage(
		connection,
		redisClient,
		userID,
	)
}
