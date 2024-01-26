package main

import (
	"context"
	"errors"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/go-redis/redis/v9"
	websocket "github.com/gorilla/websocket"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	mauth "github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	profiler "github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/app"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/repository"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func init() {
	go profiler.InitProfiling("chat-service", "1.0.0")
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
	// Parse query parameters
	query := r.URL.Query()
	authToken := query.Get("token")

	// Validate JWT token
	token, err := validateJwt(authToken)
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
	connection := new(Connection)
	connection.Socket = conn

	log.Println("WS: " + userID)
	go ListenMessages(connection, userID)
	//go SendMessage(connection)
}

func ListenMessages(conn *Connection, viewerID string) {
	repo := repository.NewChatRepository()
	rooms, _ := repo.GetInbox(viewerID)

	for _, room := range rooms {
		roomId := room.Room.Id
		go ListenRoom(conn, roomId)
	}
}

func ListenRoom(conn *Connection, roomId string) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-18624.c228.us-central1-1.gce.cloud.redislabs.com:18624",
		Password: "mBiW18GNIgPzQTbBMDEz71UVsAcNDOYF",
		DB:       0,
	})

	sub := client.Subscribe(context.Background(), roomId)
	ch := sub.Channel()
	for {
		select {
		case msg := <-ch:
			log.Println("Received msg from redis pubsub")
			var chatMessage chat.Message
			err := protojson.Unmarshal([]byte(msg.Payload), &chatMessage)
			chatMessage.Status = chat.Message_DELIVERED
			bytes, err := protojson.Marshal(&chatMessage)

			err = conn.Send(bytes)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

//func SendMessage(conn *Connection) {
//	repo := repository.NewChatRepository()
//
//	for {
//		var chatMessage chat.Message
//		_, b, err := conn.Socket.ReadMessage()
//		if err != nil {
//			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
//				log.Printf("error: %v", err)
//			}
//			break
//		}
//
//		err = proto.Unmarshal(b, &chatMessage)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//
//		msg, err := repo.SendMessage(chatMessage.SenderId, chatMessage.RoomId, chatMessage.Text)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		// Acknowledge message
//		msg.Status = chat.Message_SENT
//		bytes, err := protojson.Marshal(msg)
//		err = conn.Send(bytes)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//}

type Connection struct {
	Socket *websocket.Conn
	mu     sync.Mutex
}

// Concurrency handling - sending messages
func (c *Connection) Send(data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Socket.WriteMessage(websocket.TextMessage, data)
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

func validateJwt(token string) (*auth.Token, error) {
	// Initialize Firebase app
	ctx := context.Background()
	app := utils.InitializeAppDefault()

	// Get auth client
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error getting Auth client: %v\n", err)
		return nil, err
	}

	// Verify ID token
	idToken, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		log.Printf("Error verifying ID token: %v\n", err)
		return nil, err
	}

	// Check if token is valid
	if idToken.Expires < time.Now().Unix() {
		log.Printf("Token has expired\n")
		return nil, err
	}

	// Token is valid
	return idToken, nil
}
