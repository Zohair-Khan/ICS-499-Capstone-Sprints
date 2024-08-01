package models

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	username string
	password string
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) ValidateProvider(username string, password string) error {
	statement := `select hashedPassword from User where username = ?`

	row := u.DB.QueryRow(statement, username)

	var hashedPassword string

	err := row.Scan(&hashedPassword)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func (u *UserModel) GetID(username string) (int, error) {
	statement := `select id from User where User.username = ?`

	row := u.DB.QueryRow(statement, username)

	var id int

	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
func (u *UserModel) GetName(id int) (string, error) {
	statement := `select fname, lname from User where User.id = ?`

	row := u.DB.QueryRow(statement, id)

	var fname string
	var lname string
	err := row.Scan(&fname, &lname)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", fname, lname), nil
}

func (*UserModel) ValidateAdmin() (bool, error) {

	return true, nil
}
