package repositories

import (
	"log"
	"time"

	"go-architecture-proto/entities/models"
	"go-architecture-proto/infrastructure"
)

// 本番用のリポジトリ
type UserRepository struct {
}

func (rep *UserRepository) GetAll() []models.User {
	// まだコネクトを都度行う
	db, err := infrastructure.ConnectDB()
	rows, err := db.Query("select id, name, email from users")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
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

func (rep *UserRepository) FindById(id string) models.User {
	var user models.User

	// コネクト
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

func (rep *UserRepository) New(user models.User) {
	// コネクト
	db, err := infrastructure.ConnectDB()

	res, err := db.Exec("INSERT INTO users(name, email, age, created_at, updated_at) VALUES(?, ?, ?, ?, ?)",
		user.Name,
		user.Email,
		user.Age,
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

func (rep *UserRepository) Update(id string, user models.User) {
	// コネクト
	db, err := infrastructure.ConnectDB()
	res, err := db.Exec("UPDATE users set name = ?, updated_at = ? WHERE id = ?", user.Name, time.Now(), id)

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

func (rep *UserRepository) Delete(id string) {
	// コネクト
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
