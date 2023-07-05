package repositories

import (
	"strconv"

	"go-architecture-proto/entities/models"
)

// モック用のリポジトリ
type MockUserRepository struct {
}

func (rep *MockUserRepository) GetAllUsers() []models.User {
	var users []models.User
	user1 := models.User{
		Id:   1,
		Name: "test",
	}
	user2 := models.User{
		Id:   2,
		Name: "test2",
	}
	users = append(users, user1)
	users = append(users, user2)

	return users
}

func (rep *MockUserRepository) FindUserById(id string) models.User {
	userId, _ := strconv.Atoi(id)
	var user = models.User{
		Id:   userId,
		Name: "test",
	}

	return user
}
