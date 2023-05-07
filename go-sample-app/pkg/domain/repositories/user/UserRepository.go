package repositories

import (
	"log"
	"time"

	"go-sample-app/pkg/domain/models"
	"go-sample-app/pkg/infrastructure"
)

// 本番用のリポジトリ
type UserRepository struct {
}

func (rep *UserRepository) GetAllUsers() []models.User {
	// まだコネクトを都度行う
	db, err := infrastructure.ConnectDB()
	rows, err := db.Query("select id, name from users")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ユーザー一覧を取得しました")
	return users
}

func (rep *UserRepository) FindUserById(id string) models.User {
	var user models.User

	// まだコネクトを都度行う
	db, err := infrastructure.ConnectDB()
	rows, err := db.Query("select id, name from users where ID = ?", id)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ユーザーを取得しました")
	log.Printf("id = %d", user.Id)
	return user
}

func (rep *UserRepository) NewUser() {
	// まだコネクトを都度行う
	db, err := infrastructure.ConnectDB()

	// TODO: 汎用的にする
	res, err := db.Exec("INSERT INTO users(name, created_at, updated_at) VALUES(?, ?, ?)",
		"name",
		time.Now(), // 現在時刻
		time.Now(), // 現在時刻
	)

	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("新規ユーザーを作成しました")
	log.Printf("id = %d, affected = %d\n", lastId, rowCnt)
}

// TODO: 汎用的にする
func (rep *UserRepository) UpdateUser(id string) {
	// まだコネクトを都度行う
	db, err := infrastructure.ConnectDB()
	// TODO: 汎用的にする
	res, err := db.Exec("UPDATE users set updated_at = ? WHERE id = ?", time.Now(), id)

	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ユーザーを更新しました")
	log.Printf("id = %s, affected = %d\n", id, rowCnt)
}

func (rep *UserRepository) DeleteUser(id string) {
	// まだコネクトを都度行う
	db, err := infrastructure.ConnectDB()
	res, err := db.Exec("DELETE FROM users where id = ?", id)

	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ユーザーを削除しました")
	log.Printf("id = %s, affected = %d\n", id, rowCnt)
}
