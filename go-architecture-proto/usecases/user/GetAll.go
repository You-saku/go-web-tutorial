package usecases

import (
	"go-architecture-proto/entities/models"
	userRepository "go-architecture-proto/entities/repositories/user"
	userService "go-architecture-proto/services/user"
)

func ShowUsers() []models.User {
	ur := userRepository.UserRepository{}
	us := userService.NewUserService(&ur)

	return us.ShowUsers()
}
