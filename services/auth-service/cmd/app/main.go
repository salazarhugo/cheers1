package main

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/auth-service/internal/app"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"time"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout
}

func main() {
	if os.Getenv("DISABLE_PROFILER") == "" {
		log.Info("Profiling enabled.")
		go profiler.InitProfiling("auth-service", "1.0.0")
	} else {
		log.Info("Profiling disabled.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Info("Defaulting to port %s", port)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Feeden: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := app.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, server)
	grpc_health_v1.RegisterHealthServer(grpcServer, server)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
