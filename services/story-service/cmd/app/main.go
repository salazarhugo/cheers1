package main

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/story/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/story-service/internal/app"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os"
)

var log *logrus.Logger

func init() {
	log = utils.InitLogrus()
}

func main() {
	if os.Getenv("DISABLE_PROFILER") == "" {
		log.Info("Profiling enabled.")
		go profiler.InitProfiling("story-service", "1.0.0")
	} else {
		log.Info("Profiling disabled.")
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
		grpc.StreamInterceptor(auth.StreamInterceptor),
	)

	server := app.NewServer()

	pb.RegisterStoryServiceServer(grpcServer, server)
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

	log.Infof("Story Service listening on port %s", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
