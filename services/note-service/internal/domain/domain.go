package domain

import (
	"github.com/salazarhugo/cheers1/services/note-service/internal/repository"
)

type Domain interface {
	CreateNoteUseCase(
		params CreateNoteUseCaseParams,
	) (string, error)
}

type domain struct {
	repository repository.Repository
}

func NewDomain() Domain {
	return &domain{
		repository: repository.NewRepository(),
	}
}
