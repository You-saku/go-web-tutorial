package usecases

import (
	"web/entities/models"
	userRepository "web/entities/repositories/user"
)

type UserUsecase struct {
	ur userRepository.IUserRepository
}

/**
 * コンストラクタ(Effective Goにあった)
 * https://d-tsuji.github.io/effective_go/documents/effective_go_ja.html#id19
 */
func NewUserUsecase(ur userRepository.IUserRepository) *UserUsecase {
	userUsecase := &UserUsecase{ur: ur}
	return userUsecase
}

func (us *UserUsecase) CreateUser(user models.User) {
	us.ur.New(user)
}

func (us *UserUsecase) GetUser(userId string) models.User {
	return us.ur.FindById(userId)
}

func (us *UserUsecase) GetUsers() []models.User {
	return us.ur.GetAll()
}

func (us *UserUsecase) UpdateUser(userId string, user models.User) {
	us.ur.Update(userId, user)
}

func (us *UserUsecase) DeleteUser(userId string) {
	us.ur.Delete(userId)
}
