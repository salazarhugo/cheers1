package main

import (
	"github.com/salazarhugo/cheers1/libs/auth"
	v1 "github.com/salazarhugo/cheers1/proto/out/cheers/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(0)
	log.Printf("grpc-user: starting server...")

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
	)

	v1.RegisterMainServer(grpcServer, NewServer())
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
