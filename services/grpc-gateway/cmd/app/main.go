package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
)

func main() {
	mux := runtime.NewServeMux()

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return
	}

	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	err = party.RegisterPartyServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8080",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(transportCredentials),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Creating a normal HTTP server
	server := http.Server{
		Handler: Cors(mux),
	}

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
