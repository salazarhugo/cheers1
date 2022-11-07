package main

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/party-service/internal/app"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
)

var log *logrus.Logger

func init() {
	log = utils.InitLogrus()
}

func main() {
	if os.Getenv("DISABLE_PROFILER") == "" {
		log.Info("Profiling enabled.")
		go profiler.InitProfiling("party-service", "1.0.0")
	} else {
		log.Info("Profiling disabled.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "9080"
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

	pb.RegisterPartyServiceServer(grpcServer, server)
	grpc_health_v1.RegisterHealthServer(grpcServer, server)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
