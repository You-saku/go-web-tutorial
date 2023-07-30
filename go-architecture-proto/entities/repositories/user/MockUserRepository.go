package repositories

import (
	"strconv"

	"go-architecture-proto/entities/models"
)

// モック用のリポジトリ
type MockUserRepository struct {
}

func (rep *MockUserRepository) GetAll() []models.User {
	var users []models.User
	user1 := models.User{
		Id:    1,
		Name:  "test",
		Email: "sample@example.com",
	}
	user2 := models.User{
		Id:    2,
		Name:  "test2",
		Email: "sample2@example.com",
	}
	users = append(users, user1)
	users = append(users, user2)

	return users
}

func (rep *MockUserRepository) FindById(id string) models.User {
	userId, _ := strconv.Atoi(id)
	var user = models.User{
		Id:    userId,
		Name:  "test",
		Email: "sample3@example.com",
	}

	return user
}

func (rep *MockUserRepository) New(user models.User) {

}

func (rep *MockUserRepository) Update(id string, user models.User) {
}

func (rep *MockUserRepository) Delete(id string) {

}
