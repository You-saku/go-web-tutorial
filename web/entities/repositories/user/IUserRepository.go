package repositories

import (
	"web/entities/models"
)

type IUserRepository interface {
	FindById(id string) models.User
	GetAll() []models.User
	New(user models.User)
	Update(id string, user models.User)
	Delete(id string)
}
