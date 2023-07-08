package usecases

import (
	userRepository "go-architecture-proto/entities/repositories/user"
	userService "go-architecture-proto/services/user"
)

func UpdateUser(userId string) {
	ur := userRepository.UserRepository{}
	us := userService.NewUserService(&ur)

	us.UpdateUser(userId)
}
