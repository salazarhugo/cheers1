package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteLikeComment(
	ctx context.Context,
	request *comment.DeleteLikeCommentRequest,
) (*comment.DeleteLikeCommentResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	commentID := request.GetCommentId()

	com, err := s.commentRepository.GetComment(commentID)
	if err != nil {
		return nil, err
	}

	err = s.commentRepository.UnLikeComment(userID, com.Id)
	if err != nil {
		return nil, err
	}

	return &comment.DeleteLikeCommentResponse{}, nil
}
