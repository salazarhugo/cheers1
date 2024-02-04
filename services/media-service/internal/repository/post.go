package repository

type MediaRepository interface {
}

type mediaRepository struct {
}

func NewPostRepository() MediaRepository {
	return &mediaRepository{}
}
