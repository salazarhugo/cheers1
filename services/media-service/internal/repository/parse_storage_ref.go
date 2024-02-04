package repository

import (
	"net/url"
	"strings"
)

func ParseStorageRef(
	storageRef string,
) (string, string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(storageRef)
	if err != nil {
		return "", "", err
	}

	// Extract the bucket name and file path
	bucketName := parsedURL.Host
	filePath := strings.TrimPrefix(parsedURL.Path, "/")

	return bucketName, filePath, nil
}
