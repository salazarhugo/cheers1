package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
)

func (s *Server) CreateComment(
	ctx context.Context,
	request *comment.CreateCommentRequest,
) (*comment.CreateCommentResponse, error) {

	err := s.commentRepository.CreateComment(request.Comment, request.PostId)
	if err != nil {
		return nil, err
	}

	return &comment.CreateCommentResponse{}, nil
}
