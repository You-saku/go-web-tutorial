package services

import (
	"go-sample-app/pkg/domain/models"
	UserRepositories "go-sample-app/pkg/domain/repositories/user"
)

// TODO: Interfaceを使う
type UserService struct {
	Rep UserRepositories.UserRepository
}

func (s *UserService) ShowUser(userId string) models.User {
	return s.Rep.FindUserById(userId)
}
