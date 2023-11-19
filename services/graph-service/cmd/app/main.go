package main

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/libs/utils/auth"
	"github.com/salazarhugo/cheers1/libs/utils/profiler"
	"github.com/salazarhugo/cheers1/services/graph-service/internal/app"
	"github.com/salazarhugo/cheers1/services/graph-service/internal/app/events"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os"
)

func init() {
	go profiler.InitProfiling("graph-service", "1.0.0")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Infof("Defaulting to port %s", port)
	}

	httpMux := http.NewServeMux()
	log.Info(httpMux)
	httpMux.HandleFunc("/claim-sub", events.ClaimSub)
	httpMux.HandleFunc("/friendship-sub", events.FriendShipSub)
	httpMux.HandleFunc("/auth-sub", events.AuthSub)

	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(auth.AuthInterceptor),
	)

	server := app.NewServer()

	pb.RegisterUserServiceServer(grpcS, server)
	grpc_health_v1.RegisterHealthServer(grpcS, server)

	mixedHandler := utils.NewHttpAndGrpcMux(httpMux, grpcS)
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
