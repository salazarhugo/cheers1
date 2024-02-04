package main

import (
	"errors"
	"fmt"
	"github.com/salazarhugo/cheers1/gen/go/cheers/media/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/logger"
	"github.com/salazarhugo/cheers1/services/media-service/internal/app"
	"github.com/salazarhugo/cheers1/services/media-service/internal/app/events"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os"
)

var log *logrus.Logger

func init() {
	log = logger.InitLogrus()
	go profiler.InitProfiling("media-service", "1.0.0")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Info("Defaulting to port %s", port)
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/post-sub", events.PostSub)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := app.NewServer()

	grpc_health_v1.RegisterHealthServer(grpcServer, server)
	media.RegisterMediaServiceServer(grpcServer, server)

	mixedHandler := utils.NewHttpAndGrpcMux(httpMux, grpcServer)
	http2Server := &http2.Server{}
	http1Server := &http.Server{
		Handler: h2c.NewHandler(mixedHandler, http2Server),
	}

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
