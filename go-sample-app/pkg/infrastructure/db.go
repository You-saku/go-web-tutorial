package infrastructure

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// エラーハンドリング
	}
	c := mysql.Config{
		DBName:               "develop",
		User:                 "user",
		Passwd:               "secret",
		Addr:                 "db:3306",
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	//db, err := sql.Open("mysql", "user:secret@tcp(db:3306)/develop") // 旧書き方

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功！")
	}

	return db, err
	// 下記は一旦コメントアウト
	// See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)
}
