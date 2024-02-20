package repository

import (
	pb "github.com/salazarhugo/cheers1/gen/go/cheers/post/v1"
	"github.com/salazarhugo/cheers1/services/post-service/internal/services"
)

func (p *postRepository) ListMapPost(
	viewerID string,
	request *pb.ListMapPostRequest,
) (*pb.ListMapPostResponse, error) {
	friendIds, err := services.GetFriendList(viewerID)
	if err != nil {
		return nil, err
	}

	posts, err := p.FeedPost(
		append(friendIds, viewerID),
		int(request.Page),
		int(request.PageSize),
	)
	if err != nil {
		return nil, err
	}

	return &pb.ListMapPostResponse{
		Posts: posts.GetPosts(),
	}, nil
}
