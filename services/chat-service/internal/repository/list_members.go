package repository

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	userpb "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
)

func (c chatRepository) ListMembers(
	context context.Context,
	request *pb.ListMembersRequest,
) ([]*user.UserItem, error) {

	membersIDs := c.cache.GetRoomMembers(request.RoomId)

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	transportCredentials := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	conn, err := grpc.DialContext(context, "android-gateway-clzdlli7.nw.gateway.dev:443",
		grpc.WithTransportCredentials(transportCredentials),
	)
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	md, ok := metadata.FromIncomingContext(context)
	if !ok {
		log.Println("not ok")
	}
	log.Println(md)
	ctx := metadata.AppendToOutgoingContext(context, "authorization", md.Get("x-forwarded-authorization")[0])

	response, err := client.GetUserItemsIn(ctx, &userpb.GetUserItemsInRequest{UserIds: membersIDs})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return response.Users, nil
}
