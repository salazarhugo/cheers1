package app

import (
	"context"
	"github.com/salazarhugo/cheers1/gen/go/cheers/media/v1"
	"github.com/salazarhugo/cheers1/services/media-service/internal/repository"
	"strconv"
)

func (s *Server) UploadMedia(
	ctx context.Context,
	request *media.UploadMediaRequest,
) (*media.UploadMediaResponse, error) {
	fileBytes := request.GetChunk()
	bucketName := "cheers-posts"

	err := repository.UploadToCloudStorage(ctx, bucketName, strconv.FormatInt(request.GetUploadId(), 10), fileBytes)
	if err != nil {
		return nil, err
	}

	return &media.UploadMediaResponse{
		UploadId: request.GetUploadId(),
	}, nil
}
