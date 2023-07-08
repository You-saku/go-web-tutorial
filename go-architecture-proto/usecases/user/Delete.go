package usecases

import (
	UserRepositories "go-architecture-proto/entities/repositories/user"
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

func (s *UserService) DeleteUser(userId string) {
	s.Rep.Delete(userId)
}
