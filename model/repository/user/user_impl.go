package user

import (
	"database/sql"
	"errors"
	"time"

	"github.com/sb-67-go-quiz-salman-input-data-buku/database/connection"
	"github.com/sb-67-go-quiz-salman-input-data-buku/structs"
	"golang.org/x/crypto/bcrypt"
)

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (u *UserRepository) RegisterUser(user structs.AuthRequest) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `Insert into users (username,password,created_by) values ($1,$2,$3,$4)`

	res, err := connection.DBConn.Exec(query, user.Username, string(password), "admin")
	if err != nil {
		return errors.New("gagal menjalankan perintah tambah data")
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal menyimpan data user")
	}

	return nil
}

func (u *UserRepository) LoginUser(username string, password string) (any, error) {
	var dataUser structs.User
	query := `Select id, password from users where username = $1`

	err := connection.DBConn.QueryRow(query, username).Scan(&dataUser.Id, &dataUser.Password)
	if err == sql.ErrNoRows {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password tidak cocok dengan username ini")
	}

	return dataUser.Id, nil
}

func (u *UserRepository) UpdateUser(id string, user structs.AuthRequest, responsibleUser string) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `Update users set username=$1, password=$2, modified_at=$3, modified_by=$4 where id=$5`
	res, err := connection.DBConn.Exec(query, user.Username, password, time.Now(), responsibleUser, id)

	if err != nil {
		return errors.New("gagal menjalankan perintah update data")
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("gagal update user, id tidak ditemukan")
	}

	return nil
}
