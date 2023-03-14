package main

import (
	"errors"
	"fmt"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/libs/auth"
	"github.com/salazarhugo/cheers1/libs/profiler"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/app"
	"github.com/salazarhugo/cheers1/services/payment-service/internal/app/events"
	"github.com/stripe/stripe-go/v72"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func init() {
	secret := os.Getenv("STRIPE_SK")
	stripe.Key = secret
	log.SetFlags(0)
	go profiler.InitProfiling("payment-service", "1.0.0")

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/auth-sub", events.AuthSub)
	httpMux.HandleFunc("/handleStripeEvents", events.HandleStripeEvent)
	//httpMux.HandleFunc("/ticket/validate", app.ValidateTicket)

	grpcS := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryInterceptor),
	)

	server := app.NewServer()

	grpc_health_v1.RegisterHealthServer(grpcS, server)
	pb.RegisterPaymentServiceServer(grpcS, server)

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
