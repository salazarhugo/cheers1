package utils

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func Serve(
	port string,
	grpcS *grpc.Server,
) error {
	httpMux := http.NewServeMux()
	mixedHandler := NewHttpAndGrpcMux(httpMux, grpcS)
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

	return err
}
