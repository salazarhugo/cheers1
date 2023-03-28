package main

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/logger"
	"github.com/salazarhugo/cheers1/services/note-service/internal/app"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"os"
)

var Logger *logrus.Logger

func init() {
	go profiler.InitProfiling(os.Getenv("K_SERVICE"), "1.0.0")
	Logger = logger.InitLogrus()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		Logger.Infof("Defaulting to port %s", port)
	}

	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := app.NewServer(Logger)

	pb.RegisterNoteServiceServer(grpcS, server)
	grpc_health_v1.RegisterHealthServer(grpcS, server)

	err := utils.Serve(port, grpcS)
	if err != nil {
		Logger.Println(err)
		return
	}
}
