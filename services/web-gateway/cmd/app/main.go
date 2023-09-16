package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/salazarhugo/cheers1/gen/go/cheers/account/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/activity/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/auth/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/drink/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/location/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/note/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/notification/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/order/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/party/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/payment/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/story/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/ticket/v1"
	"github.com/salazarhugo/cheers1/gen/go/cheers/user/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/grpc-gateway/internal"
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
				_, err := client.VerifyIDToken(ctx, jwt)
				if err != nil {
					log.Printf("error verifying ID token: %v\n", err)
				}
			}

			jwtPayload := strings.Split(jwt, ".")[1]

			md := metadata.New(
				map[string]string{
					"x-apigateway-api-userinfo": jwtPayload,
				},
			)

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

	if err != nil {
		log.Println(err)
	}

	ctx := context.Background()
	isProd := flag.Bool("prod", true, "True if its in a production environment ")
	flag.Parse()

	options := []grpc.DialOption{
		grpc.WithUnaryInterceptor(internal.CloudRunInterceptor),
	}

	if *isProd {
		log.Println("Running in production environment")
	}

	options = append(options, grpc.WithTransportCredentials(transportCredentials))

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
	err = location.RegisterLocationServiceHandlerFromEndpoint(ctx, mux, "location-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = drink.RegisterDrinkServiceHandlerFromEndpoint(ctx, mux, "drink-service-r3a2dr4u4a-nw.a.run.app:443", options)
	err = note.RegisterNoteServiceHandlerFromEndpoint(ctx, mux, "note-service-r3a2dr4u4a-nw.a.run.app:443", options)

	handler := internal.ChainMiddleware(mux,
		internal.Cors,
		internal.LoggerMiddleware,
		//internal.BlockEndpointMiddleware,
	)

	// Creating a normal HTTP server
	server := http.Server{
		Handler: handler,
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
