package controllers

import (
	"net/http"

	repository "go-sample-app/pkg/domain/repositories/todo"
	service "go-sample-app/pkg/domain/services/todo"
)

// POST:posts/
func TodoCreateHandler(w http.ResponseWriter, r *http.Request) {

	requestMethod := r.Method // これでhttpリクエストのメソッドを取得
	if requestMethod != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 本番用リポジトリ層
	repository := repository.TodoRepository{}
	// サービス層作成
	service := service.NewTodoService(&repository)

	service.CreateTodo()
}
