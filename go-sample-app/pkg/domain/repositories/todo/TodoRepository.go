package repositories

import (
	"fmt"
)

// 本番用のリポジトリ
type TodoRepository struct {
}

func (rep *TodoRepository) NewTodo() {
	// TODO: todoを作成する
	fmt.Println("create a todo.")
}
