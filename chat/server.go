package main

import (
	"chat/chatpb"
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	glog "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var grpcLog glog.LoggerV2

func init() {
	//cors.j
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func main() {
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

func streamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {

	md, err := authorize(ss.Context())
	if err != nil {
		return err
	}

	err = ss.SetHeader(md)
	ss.SendHeader(md)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Set headers stream", md)

	return handler(srv, ss)
}

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	start := time.Now()

	md, err := authorize(ctx)
	if err != nil {
		return req, err
	}

	err = grpc.SetHeader(ctx, md)
	if err != nil {
		return nil, err
	}

	h, err := handler(ctx, req)

	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

// authorize function authorizes the token received from Metadata
func authorize(ctx context.Context) (metadata.MD, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return md, status.Errorf(codes.InvalidArgument, "Failed retrieving metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return md, status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}
	token := authHeader[0]

	// validateToken function validates the token
	validToken, err := validateToken(token)
	if err != nil {
		return md, status.Errorf(codes.Unauthenticated, err.Error())
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "userId", validToken.UID)

	return metadata.Pairs("userId", validToken.UID), nil
}

func validateToken(idToken string) (*auth.Token, error) {
	app := initializeAppDefault()
	client := initializeAuthClient(app)

	idToken = strings.Split(idToken, "Bearer ")[1]

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		log.Println("Invalid Authorization Token", idToken)
		return nil, err
	}
	return token, nil
}
