package main

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/party-service/internal/app"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os"
	"strings"
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
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	m := cmux.New(listener)

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", app.PaymentSub)

	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := app.NewServer()

	pb.RegisterPartyServiceServer(grpcS, server)
	grpc_health_v1.RegisterHealthServer(grpcS, server)

	httpS := http.Server{
		Handler: httpMux,
	}

	httpL := m.Match(cmux.HTTP1Fast())
	grpcL := m.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"),
	)

	go func() {
		err := httpS.Serve(httpL)
		if err != nil {
			log.Println(err)
		}
	}()
	go func() {
		err := grpcS.Serve(grpcL)
		if err != nil {
			log.Println(err)
		}
	}()

	// Start cmux serving.
	if err := m.Serve(); !strings.Contains(err.Error(),
		"use of closed network connection") {
		log.Fatal(err)
	}
}
