package repositories

import (
	"go-sample-app/pkg/domain/models"
)

// TODO: これをうまく使いたい
type IUserRepository interface {
	FindUserById(id string) models.User
}
