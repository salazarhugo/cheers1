package main

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/salazarhugo/cheers1/genproto/grpc/health/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	pb "github.com/salazarhugo/cheers1/services/postservice/genproto/cheers/post/v1"
	"github.com/salazarhugo/cheers1/services/postservice/internal/app"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

func init() {
	log.SetFlags(0)
}

func main() {
	if os.Getenv("DISABLE_PROFILER") == "" {
		log.Print("Profiling enabled.")
		go profiler.InitProfiling("postservice", "1.0.0")
	} else {
		log.Print("Profiling disabled.")
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
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := app.NewServer()

	pb.RegisterPostServiceServer(grpcServer, server)
	health.RegisterHealthServer(grpcServer, server)

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

	log.Printf("http server listening at %v", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
