package domain

import "github.com/salazarhugo/cheers1/services/user-service/internal/repository"

// UserService is responsible for implementing the business logic for user management.
type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance of the UserService.
func NewUserService() *UserService {
	return &UserService{userRepository: repository.NewUserRepository()}
}
