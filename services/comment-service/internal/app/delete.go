package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) DeleteComment(
	ctx context.Context,
	request *comment.DeleteCommentRequest,
) (*comment.DeleteCommentResponse, error) {
	userID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	com, err := s.commentRepository.GetComment(request.CommentId)
	if err != nil {
		return nil, err
	}

	if com.UserId != userID {
		return nil, status.Error(codes.PermissionDenied, "you can't delete someone else's comment")
	}

	err = s.commentRepository.DeleteComment(
		request.CommentId,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &comment.DeleteCommentResponse{}, nil
}
