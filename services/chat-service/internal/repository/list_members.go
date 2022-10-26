package repository

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	userpb "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func (c chatRepository) ListMembers(
	context context.Context,
	request *pb.ListMembersRequest,
) ([]*user.UserItem, error) {

	membersIDs := c.cache.GetRoomMembers(request.RoomId)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("android-gateway-clzdlli7.nw.gateway.dev:443", opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	md, ok := metadata.FromIncomingContext(context)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Failed retrieving metadata")
	}
	log.Println(md)
	log.Println("mmh...")

	response, err := client.GetUserItemsIn(context, &userpb.GetUserItemsInRequest{UserIds: membersIDs})
	if err != nil {
		return nil, err
	}

	return response.Users, nil
}
