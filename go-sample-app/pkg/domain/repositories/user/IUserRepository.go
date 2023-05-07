package repositories

import (
	"go-sample-app/pkg/domain/models"
)

type IUserRepository interface {
	FindById(id string) models.User
	GetAll() []models.User
	New()
	Update(id string)
	Delete(id string)
}
