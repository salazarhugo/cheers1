package main

import (
	grpcweb "github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/activity/v1"
	auth "github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	app "github.com/salazarhugo/cheers1/services/activity-service/internal/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	if os.Getenv("DISABLE_PROFILER") == "" {
		log.Println("Profiling enabled.")
		go profiler.InitProfiling("post-service", "1.0.0")
	} else {
		log.Println("Profiling disabled.")
	}

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	StartGrpcServer(port)
}

func StartGrpcServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Feeden: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := app.NewServer()

	pb.RegisterActivityServiceServer(grpcServer, server)
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
	log.Printf("Post Service Feedening on port %s", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
