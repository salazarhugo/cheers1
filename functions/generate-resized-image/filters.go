package generate_resized_image

import "strings"

func (s StorageObjectData) ShouldResize() bool {
	contentType := s.ContentType

	if !strings.Contains(contentType, "image/") {
		return false
	}

	return true
}
