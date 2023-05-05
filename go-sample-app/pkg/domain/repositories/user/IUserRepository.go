package repositories

import (
	"go-sample-app/pkg/domain/models"
)

type IUserRepository interface {
	FindUserById(id string) models.User
}
