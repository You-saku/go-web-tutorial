package services

import (
	"go-sample-app/pkg/domain/models"
	UserRepositories "go-sample-app/pkg/domain/repositories/user"
)

type UserService struct {
	Rep UserRepositories.IUserRepository
}

/**
 * コンストラクタ(Effective Goにあった)
 * http://go.shibu.jp/effective_go.html#id16
 */
func NewUserService(rep UserRepositories.IUserRepository) *UserService {
	userService := new(UserService)
	userService.Rep = rep

	return userService
}

func (s *UserService) ShowUser(userId string) models.User {
	return s.Rep.FindUserById(userId)
}
