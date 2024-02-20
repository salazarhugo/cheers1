package app

import (
	"context"
	"fmt"
	"github.com/salazarhugo/cheers1/gen/go/cheers/media/v1"
	"github.com/salazarhugo/cheers1/services/media-service/internal/models"
	"github.com/salazarhugo/cheers1/services/media-service/internal/repository"
	"strconv"
)

func (s *Server) UploadMedia(
	ctx context.Context,
	request *media.UploadMediaRequest,
) (*media.UploadMediaResponse, error) {
	fileBytes := request.GetChunk()
	bucketName := "cheers-posts"
	objectName := strconv.FormatInt(request.GetUploadId(), 10)

	object, err := repository.UploadImageToCloudStorage(ctx, bucketName, objectName, fileBytes)
	if err != nil {
		return nil, err
	}

	objectAttributes, err := object.Attrs(ctx)
	if err != nil {
		return nil, err
	}

	mediaID, err := s.mediaRepository.CreateMedia(&models.Media{
		Url:  objectAttributes.MediaLink,
		Ref:  fmt.Sprintf("gs://%s/%s", bucketName, objectName),
		Type: objectAttributes.ContentType,
	})
	if err != nil {
		return nil, err
	}

	return &media.UploadMediaResponse{
		MediaId: mediaID,
	}, nil
}
