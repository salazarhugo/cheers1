package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"log"
)

func (s *Server) ListComment(
	ctx context.Context,
	request *comment.ListCommentRequest,
) (*comment.ListCommentResponse, error) {

	items, err := s.commentRepository.ListComment(request.PostId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &comment.ListCommentResponse{
		Items: items,
	}, nil
}
