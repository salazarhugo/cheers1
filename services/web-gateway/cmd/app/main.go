package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/felixge/httpsnoop"
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
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
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
		grpc.WithUnaryInterceptor(clientInterceptor),
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

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	start := time.Now()
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Failed retrieving metadata")
	}

	audience, err := getAudience(cc.Target())
	if err != nil {
		log.Println(err)
	}

	accessToken, err := generateAccessToken(ctx, audience)
	if err != nil {
		log.Println(err)
	}

	log.Println(audience, accessToken)

	md.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	md.Set("X-Serverless-Authorization", fmt.Sprintf("Bearer %s", accessToken))

	ctx = metadata.NewOutgoingContext(ctx, md)

	// Calls the invoker to execute RPC
	err = invoker(ctx, method, req, reply, cc, opts...)

	// Logic after invoking the invoker
	log.Printf("Invoked RPC method=%s; Duration=%s; Error=%v", method, time.Since(start), err)

	return err
}

func getAudience(target string) (string, error) {
	host, _, err := net.SplitHostPort(target)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s", host), nil
}

func generateAccessToken(ctx context.Context, audience string) (string, error) {
	tokenSource, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		return "", err
	}

	token, err := tokenSource.Token()
	if err != nil {
		log.Println(err)
	}

	return token.AccessToken, nil
}
