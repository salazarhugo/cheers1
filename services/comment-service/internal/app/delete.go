package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/comment/v1"
	"log"
)

func (s *Server) DeleteComment(
	ctx context.Context,
	request *comment.DeleteCommentRequest,
) (*comment.DeleteCommentResponse, error) {
	//userID, err := utils.GetUserId(ctx)
	//if err != nil {
	//	return nil, status.Error(codes.Internal, "failed to retrieve userID")
	//}

	err := s.commentRepository.DeleteComment(
		request.PostId,
		request.CommentId,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &comment.DeleteCommentResponse{}, nil
}
