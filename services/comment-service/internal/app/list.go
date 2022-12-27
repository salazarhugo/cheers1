package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
)

func (s *Server) ListComment(
	ctx context.Context,
	request *comment.ListCommentRequest,
) (*comment.ListCommentResponse, error) {

	comments, err := s.commentRepository.ListComment(request.PostId)
	if err != nil {
		return nil, err
	}

	return &comment.ListCommentResponse{
		Comments: comments,
	}, nil
}
