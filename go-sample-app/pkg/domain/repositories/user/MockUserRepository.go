package repositories

import (
	"strconv"

	"go-sample-app/pkg/domain/models"
)

// モック用のリポジトリ
type MockUserRepository struct {
}

func (rep *MockUserRepository) FindUserById(id string) models.User {
	userId, _ := strconv.Atoi(id)
	var user = models.User{
		Id:   userId,
		Name: "test",
	}

	return user
}
