package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"go-sample-app/models"
)

func connectDB() {
	// db
	db, err := sql.Open("mysql", "user:secret@tcp(db:3306)/develop") // ひとまずベタ書き
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
		return
	} else {
		fmt.Println("データベース接続成功！")
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

// routing
func sampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func dbShowHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	userId := strings.TrimPrefix(r.URL.Path, "/users/")

	// コネクトを都度行う感じ
	db, err := sql.Open("mysql", "user:secret@tcp(db:3306)/develop") // ひとまずベタ書き
	rows, err := db.Query("select id, name from users where ID = ?", userId)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user.Id, user.Name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "This is %s' data.", user.Name)
}

func main() {
	connectDB()

	// net/http
	http.HandleFunc("/", sampleHandler)
	http.HandleFunc("/users/", dbShowHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
