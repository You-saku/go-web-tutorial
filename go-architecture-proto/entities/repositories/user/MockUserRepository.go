package repositories

import (
	"database/sql"
	"strconv"

	"go-architecture-proto/entities/models"
)

// モック用のリポジトリ
type MockUserRepository struct {
}

func (rep *MockUserRepository) GetAll() []models.User {
	var users []models.User
	user1 := models.User{
		Id:   1,
		Name: "test",
		Email: sql.NullString{
			String: "sample@example.com",
			Valid:  true,
		},
	}
	user2 := models.User{
		Id:   2,
		Name: "test2",
		Email: sql.NullString{
			String: "sample2@example.com",
			Valid:  true,
		},
	}
	users = append(users, user1)
	users = append(users, user2)

	return users
}

func (rep *MockUserRepository) FindById(id string) models.User {
	userId, _ := strconv.Atoi(id)
	var user = models.User{
		Id:   userId,
		Name: "test",
		Email: sql.NullString{
			String: "sample3@example.com",
			Valid:  true,
		},
	}

	return user
}

func (rep *MockUserRepository) New(user models.User) {

}

func (rep *MockUserRepository) Update(id string, user models.User) {
}

func (rep *MockUserRepository) Delete(id string) {

}
