package usecases

import (
	"web/entities/models"
	userRepository "web/entities/repositories/user"
)

type UserUsecase struct {
	rep userRepository.IUserRepository
}

/**
 * コンストラクタ(Effective Goにあった)
 * https://d-tsuji.github.io/effective_go/documents/effective_go_ja.html#id19
 */
func NewUserUsecase(rep userRepository.IUserRepository) *UserUsecase {
	userUsecase := new(UserUsecase)
	userUsecase.rep = rep

	return userUsecase
}

func (us *UserUsecase) CreateUser(user models.User) {
	us.rep.New(user)
}

func (us *UserUsecase) GetUser(userId string) models.User {
	return us.rep.FindById(userId)
}

func (us *UserUsecase) GetUsers() []models.User {
	return us.rep.GetAll()
}

func (us *UserUsecase) UpdateUser(userId string, user models.User) {
	us.rep.Update(userId, user)
}

func (us *UserUsecase) DeleteUser(userId string) {
	us.rep.Delete(userId)
}
