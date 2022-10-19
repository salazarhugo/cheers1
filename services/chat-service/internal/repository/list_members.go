package repository

import (
	"context"
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	userpb "github.com/salazarhugo/cheers1/genproto/cheers/user/v1"
	"google.golang.org/grpc"
	"os"
)

func (c chatRepository) ListMembers(
	request *pb.ListMembersRequest,
) ([]*user.UserItem, error) {

	membersIDs := c.cache.GetRoomMembers(request.RoomId)

	var opts []grpc.DialOption
	conn, err := grpc.Dial(os.Getenv("GATEWAY_URL"), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	response, err := client.GetUser(context.Background(), &userpb.GetUserRequest{Id: membersIDs[0]})

	users := make([]*user.UserItem, 0)
	users = append(users, response.GetUser())

	return users, nil
}
