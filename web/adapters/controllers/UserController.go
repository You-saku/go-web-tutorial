package controllers

import (
	"web/entities/models"
	"encoding/json"
	"net/http"
	"strings"

	userRepository "web/entities/repositories/user"
	userUsecase "web/usecases/user"
)

type UserController struct {
	us userUsecase.UserUsecase
}

func NewUserController() *UserController {
	// 本番用リポジトリ層(interface利用)
	// ur := userRepository.MockUserRepository{} // これはモックリポジトリ
	ur := userRepository.UserRepository{}
	// サービス層作成
	us := userUsecase.NewUserUsecase(&ur)
	// UserControllerのインスタンスを返す
	return &UserController{us: *us}
}

func (uc *UserController) Create (w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user) // リクエストボディをデコード

	uc.us.CreateUser(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user) // 作成したユーザーを返す場合
	return
}

func (uc *UserController) Update (w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimPrefix(r.URL.Path, "/users/")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user) // リクエストボディをデコード

	uc.us.UpdateUser(userId, user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user) // 更新したユーザーを返す場合
	return
}

func (uc *UserController) Delete (w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimPrefix(r.URL.Path, "/users/")

	uc.us.DeleteUser(userId)
	w.WriteHeader(http.StatusOK)
	return
}

func (uc *UserController) GetById (w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimPrefix(r.URL.Path, "/users/")
	user := uc.us.GetUser(userId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return
}

func (uc *UserController) GetAll (w http.ResponseWriter, r *http.Request) {
	users := uc.us.GetUsers()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
	return
}

// 旧式
// // GET:users/{id}
// // PUT:users/{id}
// // DELETE:users/{id}
// func UserHandler(w http.ResponseWriter, r *http.Request) {
// 	// 本番用リポジトリ層
// 	ur := userRepository.UserRepository{}
// 	// サービス層作成
// 	userUsecase := userUsecase.NewUserUsecase(&ur)

// 	userId := strings.TrimPrefix(r.URL.Path, "/users/")
// 	requestMethod := r.Method // これでhttpリクエストのメソッドを取得

// 	// GET
// 	if requestMethod == "GET" {
// 		user := userUsecase.GetUser(userId)

// 		json.NewEncoder(w).Encode(user)
// 		//fmt.Fprintf(w, "id = %d name = %s email = %s\n", user.Id, user.Name, user.Email) // for debug
// 		return
// 	}
// 	// PUT
// 	if requestMethod == "PUT" {
// 		var user models.User
// 		json.NewDecoder(r.Body).Decode(&user) // リクエストボディをデコード
// 		fmt.Fprintf(w, "name = %s\n", user.Name)
// 		userUsecase.UpdateUser(userId, user)
// 		return
// 	}
// 	// DELETE
// 	if requestMethod == "DELETE" {
// 		userUsecase.DeleteUser(userId)
// 		return
// 	}

// 	// それ以外のメソッドはエラー
// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 	return
// }
