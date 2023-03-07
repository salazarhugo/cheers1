package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"log"
)

func (s *Server) ListReplies(
	ctx context.Context,
	request *comment.ListRepliesRequest,
) (*comment.ListRepliesResponse, error) {

	items, err := s.commentRepository.ListReplies(request.CommentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &comment.ListRepliesResponse{
		Items: items,
	}, nil
}
