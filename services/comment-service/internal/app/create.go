package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"github.com/salazarhugo/cheers1/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateComment(
	ctx context.Context,
	request *comment.CreateCommentRequest,
) (*comment.CreateCommentResponse, error) {
	viewerID, err := utils.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrieve userID")
	}

	var commentID string
	if request.ReplyToCommentId != "" {
		commentID, err = s.commentRepository.CreateReplyComment(
			viewerID,
			request.Comment,
			request.PostId,
			request.ReplyToCommentId,
		)
	} else {
		commentID, err = s.commentRepository.CreateComment(
			viewerID,
			request.Comment,
			request.PostId,
		)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	item, err := s.commentRepository.GetCommentItem(viewerID, commentID)
	if err != nil {
		return nil, err
	}

	return &comment.CreateCommentResponse{
		Item: item,
	}, nil
}
