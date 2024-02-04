package generate_resized_image

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/salazarhugo/cheers1/functions/generate-resized-image/repository"
	"image"
	"image/jpeg"
	"time"
)

func init() {
	functions.CloudEvent("GenerateResizedImage", generateResizedImage)
}

// StorageObjectData contains metadata of the Cloud Storage object.
type StorageObjectData struct {
	Bucket         string    `json:"bucket,omitempty"`
	Name           string    `json:"name,omitempty"`
	Metageneration int64     `json:"metageneration,string,omitempty"`
	TimeCreated    time.Time `json:"timeCreated,omitempty"`
	Updated        time.Time `json:"updated,omitempty"`
	ContentType    string    `json:"contentType,omitempty"`
}

// GenerateResizedImage /*
func generateResizedImage(ctx context.Context, e event.Event) error {
	var data StorageObjectData
	if err := e.DataAs(&data); err != nil {
		return fmt.Errorf("event.DataAs: %v", err)
	}

	if !data.ShouldResize() {
		return nil
	}

	// Download the file from Cloud Storage if needed
	bucketName := data.Bucket
	objectName := data.Name

	post, err := repository.GetPost(data.Name)
	if err != nil {
		return err
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("Bucket(%q).Object(%q).NewReader: %v", bucketName, objectName, err)
	}
	defer rc.Close()

	// Decode the image
	img, _, err := image.Decode(rc)
	if err != nil {
		return fmt.Errorf("image.Decode: %v", err)
	}

	// Resize the image to 1080x1080
	resizedImg, err := ResizeImage(img, 1080, 1080)
	if err != nil {
		return fmt.Errorf("ResizeImage: %v", err)
	}

	// Encode the resized image as JPEG
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, resizedImg, nil)
	if err != nil {
		return fmt.Errorf("jpeg.Encode: %v", err)
	}

	// Upload the resized image back to Cloud Storage
	resizedObjectName := "resized_" + objectName
	err = UploadToCloudStorage(ctx, bucketName, resizedObjectName, buf.Bytes())
	if err != nil {
		return fmt.Errorf("UploadToCloudStorage: %v", err)
	}

	return nil
}
