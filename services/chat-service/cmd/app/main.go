package main

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	profiler "github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/app"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"os"
)

func init() {
}

func main() {
	if os.Getenv("DISABLE_PROFILER") == "" {
		log.Println("Profiling enabled.")
		go profiler.InitProfiling("chat-service", "1.0.0")
	} else {
		log.Println("Profiling disabled.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		//grpc.UnaryInterceptor(auth.UnaryInterceptor),
		//grpc.StreamInterceptor(auth.StreamInterceptor),
	)

	server := app.NewServer()

	pb.RegisterChatServiceServer(grpcServer, server)
	grpc_health_v1.RegisterHealthServer(grpcServer, server)

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	srv := &http.Server{
		Handler: grpcWebServer,
		Addr:    ":8081",
	}

	log.Printf("chat Service listening on port %s", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
