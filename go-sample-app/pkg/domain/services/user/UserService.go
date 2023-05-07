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
 * https://d-tsuji.github.io/effective_go/documents/effective_go_ja.html#id19
 */
func NewUserService(rep UserRepositories.IUserRepository) *UserService {
	userService := new(UserService)
	userService.Rep = rep

	return userService
}

func (s *UserService) ShowUsers() []models.User {
	return s.Rep.GetAll()
}

func (s *UserService) ShowUser(userId string) models.User {
	return s.Rep.FindById(userId)
}

func (s *UserService) CreateUser() {
	s.Rep.New()
}

func (s *UserService) UpdateUser(userId string) {
	s.Rep.Update(userId)
}

func (s *UserService) DeleteUser(userId string) {
	s.Rep.Delete(userId)
}
