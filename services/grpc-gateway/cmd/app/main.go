package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/notification/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/story/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
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
			fields := strings.Fields(header)
			if len(fields) < 2 {
				return nil
			}
			jwt := fields[1]
			app := utils.InitializeAppDefault()
			client, err := app.Auth(ctx)
			if err != nil {
				log.Printf("error getting Auth client: %v\n", err)
			} else {
				token, err := client.VerifyIDToken(ctx, jwt)
				if err != nil {
					log.Printf("error verifying ID token: %v\n", err)
				}
				log.Printf("Verified ID token: %v\n", token)
			}

			jwtPayload := strings.Split(jwt, ".")[1]

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

	ctx := context.Background()
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(transportCredentials),
	}

	err = chat.RegisterChatServiceHandlerFromEndpoint(ctx, mux, "chat-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "user-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = party.RegisterPartyServiceHandlerFromEndpoint(ctx, mux, "party-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = post.RegisterPostServiceHandlerFromEndpoint(ctx, mux, "post-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = story.RegisterStoryServiceHandlerFromEndpoint(ctx, mux, "story-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = activity.RegisterActivityServiceHandlerFromEndpoint(ctx, mux, "activity-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = chat.RegisterChatServiceHandlerFromEndpoint(ctx, mux, "chat-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, "order-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = notification.RegisterNotificationServiceHandlerFromEndpoint(ctx, mux, "notification-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "auth-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = account.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, "account-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = payment.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, "payment-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = ticket.RegisterTicketServiceHandlerFromEndpoint(ctx, mux, "ticket-service-r3a2dr4u4a-nw.a.run.app:443", options)

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
