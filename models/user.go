package models

import (
	"database/sql"
	"fmt"
	db "music/database"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	password  string
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetUserById(id int) (user User) {
	sqlStatment := `SELECT id,username,created_at,updated_at FROM users WHERE id = $1 LIMIT 1`
	row := db.DB.QueryRow(sqlStatment, id)
	err := row.Scan(&user.Id, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	return
}

func GetUserByUsername(id int) (user User) {
	sqlStatment := `SELECT id,username,created_at,updated_at FROM users WHERE username = $1 LIMIT 1`
	row := db.DB.QueryRow(sqlStatment, id)
	err := row.Scan(&user.Id, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	return
}

func CreateUser(username string, password string) (id int, err error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return id, err
	}
	sqlStatment := `INSERT INTO users (username, password, created_at, updated_at) VALUES ($1,$2,$3,$4) RETURNING id`
	time := time.Now()
	row := db.DB.QueryRow(sqlStatment, username, hashedPassword, time, time)
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	return
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
