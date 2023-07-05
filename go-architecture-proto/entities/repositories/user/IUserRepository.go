package repositories

import (
	"go-architecture-proto/entities/models"
)

type IUserRepository interface {
	FindById(id string) models.User
	GetAll() []models.User
	New()
	Update(id string)
	Delete(id string)
}
