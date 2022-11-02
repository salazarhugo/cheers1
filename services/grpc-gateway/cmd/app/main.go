package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/libs/auth/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			// send all the headers received from the client
			jwt := strings.Fields(header)[1]

			app := utils.InitializeAppDefault()
			client, err := app.Auth(ctx)
			if err != nil {
				log.Fatalf("error getting Auth client: %v\n", err)
			}

			token, err := client.VerifyIDToken(ctx, jwt)
			if err != nil {
				log.Fatalf("error verifying ID token: %v\n", err)
			}

			jwtPayload := strings.Split(jwt, ".")[1]

			log.Printf("Verified ID token: %v\n", token)
			log.Println(jwtPayload)

			md := metadata.Pairs("x-apigateway-api-userinfo", jwtPayload)
			return md
		},
		))

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
	}

	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	err = party.RegisterPartyServiceHandlerFromEndpoint(context.Background(), mux, "party-service-r3a2dr4u4a-nw.a.run.app:443",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(transportCredentials),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Creating a normal HTTP server
	server := http.Server{
		Handler: withLogger(Cors(mux)),
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}

func withLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}
