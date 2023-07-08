package usecases

import (
	userRepository "go-architecture-proto/entities/repositories/user"
	userService "go-architecture-proto/services/user"
)

func CreateUser() {
	// 本番用リポジトリ層
	ur := userRepository.UserRepository{}
	// サービス層作成
	us := userService.NewUserService(&ur)

	us.CreateUser()
}
