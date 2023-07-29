package models

import (
	"database/sql"
)

type User struct {
	Id    int
	Name  string
	Email sql.NullString // null許容(出力時注意)
	Age   int
}

/**
 * type NullString struct {
 *	 String string
 *	 Valid  bool // Valid is true if String is not NULL
 * }
 */
