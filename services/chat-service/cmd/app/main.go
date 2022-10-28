package main

import (
	context2 "context"
	"crypto/tls"
	"crypto/x509"
	grpcweb "github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	userpb "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
	auth "github.com/salazarhugo/cheers1/libs/auth"
	profiler "github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/chat-service/internal/app"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"net/http"
	"os"
)

func init() {
}

func test() {
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return
	}

	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	context := context2.Background()
	conn, err := grpc.DialContext(context, "android-gateway-clzdlli7.nw.gateway.dev:443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	ctx := metadata.AppendToOutgoingContext(context, "authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjNmNjcyNDYxOTk4YjJiMzMyYWQ4MTY0ZTFiM2JlN2VkYTY4NDZiMzciLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vY2hlZXJzLWEyNzVlIiwiYXVkIjoiY2hlZXJzLWEyNzVlIiwiYXV0aF90aW1lIjoxNjYxODAwMDMyLCJ1c2VyX2lkIjoid29GdHBYVVQ5aVRVSU1Zc3NHNWtJSDVhbmg0MyIsInN1YiI6IndvRnRwWFVUOWlUVUlNWXNzRzVrSUg1YW5oNDMiLCJpYXQiOjE2NjY4MDYwNTYsImV4cCI6MTY2NjgwOTY1NiwiZW1haWwiOiJodWdvYnJvY2s3NCtjaGVlcnNAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsiaHVnb2Jyb2NrNzQrY2hlZXJzQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.a_sm41lK6zrGDniAvusMc3XmYIO0Oxq6e0BdHOyV-0op0bz289f0GXXYbqJhzLMIJ3KZgwinCWdXE4Saq0M_w3HrEKO4x7oimRoFMOvybody1LediPji-eYSMN1W3lIyd5jZZqZvsyQYfJa139ZkGkLkHwz2uoBdwl2dmMvdEDcqCGfZbYIjbiocH9Qy1ssiDL7dzYuNc2L0T-ZzcRzfFIGXeUAvccwZEVGkHvERp2oaR1_CFdWK6qoSpvx5KRzfwoZxJ_f6OBtlpi4Mv2ns-OB70i4ZWu20ZtxNOYVHISjVHbzrbAdl4kaOG_m3A7myvRo0xAjxXOLO6XuGTO-GKw")

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return
	}
	log.Println(md)

	response, err := client.GetUserItemsIn(ctx, &userpb.GetUserItemsInRequest{UserIds: []string{"woFtpXUT9iTUIMYssG5kIH5anh43"}})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(response)
}

func main() {
	if os.Getenv("DISABLE_PROFILER") == "" {
		log.Println("Profiling enabled.")
		go profiler.InitProfiling("chat-service", "1.0.0")
	} else {
		log.Println("Profiling disabled.")
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

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
		grpc.StreamInterceptor(auth.StreamInterceptor),
	)

	server := app.NewServer()

	pb.RegisterChatServiceServer(grpcServer, server)
	grpc_health_v1.RegisterHealthServer(grpcServer, server)

	go func() {
		if err = grpcServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	srv := &http.Server{
		Handler: grpcWebServer,
		Addr:    ":8081",
	}

	log.Printf("chat Service listening on port %s", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
