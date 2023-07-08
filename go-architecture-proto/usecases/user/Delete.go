package usecases

import (
	userRepository "go-architecture-proto/entities/repositories/user"
	userService "go-architecture-proto/services/user"
)

func DeleteUser(userId string) {
	ur := userRepository.UserRepository{}
	us := userService.NewUserService(&ur)

	us.DeleteUser(userId)
}
