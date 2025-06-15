package data

import (
	"database/sql"

	"github.com/sb-67-go-quiz-salman-input-data-buku/database/connection"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
)

func IsUserExists(user structs.AuthRequest) bool {
	sqlStatement := `SELECT id from users where username = $1 and password = $2 returning id`

	var id int
	err := connection.DBConn.QueryRow(sqlStatement, user.Username, user.Password).Scan(&id)

	return err == sql.ErrNoRows
}

func IsCategoryExists(name string) bool {
	sqlStatement := `SELECT id from categories where name = $1 returning id`

	var id int
	err := connection.DBConn.QueryRow(sqlStatement, name).Scan(&id)

	return err == sql.ErrNoRows
}

func IsBookExists(title string) bool {
	sqlStatement := `SELECT id from book where title = $1 returning id`

	var id int
	err := connection.DBConn.QueryRow(sqlStatement, title).Scan(&id)

	return err == sql.ErrNoRows
}
