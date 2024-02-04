package main

import (
	"errors"
	"fmt"
	"github.com/salazarhugo/cheers1/gen/go/cheers/notification/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/notification-service/internal/app/events"
	http3 "github.com/salazarhugo/cheers1/services/notification-service/internal/app/http"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"os"
)

func init() {
	go profiler.InitProfiling("notification-service", "1.0.0")
}

func main() {
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", events.PostSub)
	httpMux.HandleFunc("/chat", events.ChatTopic)
	httpMux.HandleFunc("/user", events.UserTopicSub)
	httpMux.HandleFunc("/friendship-sub", events.FriendShipSub)
	httpMux.HandleFunc("/comment-sub", events.CommentSub)

	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := http3.NewServer()

	grpc_health_v1.RegisterHealthServer(grpcS, server)
	notification.RegisterNotificationServiceServer(grpcS, server)

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
