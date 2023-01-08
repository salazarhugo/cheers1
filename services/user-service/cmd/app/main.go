package main

import (
	"errors"
	"fmt"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/user-service/internal/app"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var log *logrus.Logger

func init() {
	go profiler.InitProfiling("user-service", "1.0.0")
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Info("Defaulting to port %s", port)
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/claim-sub", app.ClaimSub)
	httpMux.HandleFunc("/friendship-sub", app.FriendShipSub)

	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := app.NewServer()

	pb.RegisterUserServiceServer(grpcS, server)
	grpc_health_v1.RegisterHealthServer(grpcS, server)

	mixedHandler := newHTTPandGRPCMux(httpMux, grpcS)
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

func newHTTPandGRPCMux(httpHand http.Handler, grpcHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("content-type"), "application/grpc") {
			grpcHandler.ServeHTTP(w, r)
			return
		}
		httpHand.ServeHTTP(w, r)
	})
}
