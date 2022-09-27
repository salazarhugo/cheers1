package main

import (
	"cheers.com/proto"
	v1 "cheers.com/proto/cheers/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	out.Hello("")
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

	v1.RegisterMainServer(grpcServer, newServer())
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
