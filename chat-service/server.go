package main

import (
	"chat/chatpb"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
	"log"
	"net"
	"net/http"
	"os"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func main() {
	log.SetFlags(0)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	grpcLog.Info("-- SERVER APP --")

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", port, err)
	}
	log.Printf("Listening @%v:%s", port, listen.Addr())

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	chatpb.RegisterChatServiceServer(grpcServer, newServer())
	go func() {
		log.Fatalf("failed to serve: %v", grpcServer.Serve(listen))
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
