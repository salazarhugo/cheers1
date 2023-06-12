package repository

type AuthRepository interface {
	CreateModerator(userID string) error
	CreateBusinessAccount(userID string) (string, error)
	VerifyUser(userID string) error
}

type authRepository struct {
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}
