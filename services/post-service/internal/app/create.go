package app

import (
	"context"
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"github.com/salazarhugo/cheers1/services/post-service/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePost(
	ctx context.Context,
	request *pb.CreatePostRequest,
) (*pb.PostResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed retrieving userID")
	}

	newPost := &repository.Post{
		UserID:   userID,
		Caption:  request.Caption,
		Location: request.LocationName,
	}
	if request.DrinkId > 0 {
		newPost.DrinkID = request.GetDrinkId()
	}

	postID, err := s.postRepository.CreatePost(userID, newPost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create post")
	}

	post, err := s.postRepository.GetPostItem(userID, postID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get post")
	}

	return post, nil
}
