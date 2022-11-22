package repository

type AuthRepository interface {
	CreateModerator(userID string) (string, error)
	CreateBusinessAccount(userID string) (string, error)
}

type authRepository struct {
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}
