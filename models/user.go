package models

import (
	db "music/database"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUserById(id int) (user User) {
	err := db.DB.Find(&user, id).Error
	if err != nil {
		panic(err)
	}
	return
}

func GetUserByUsername(username string) (user User) {
	err := db.DB.Where("username = ?", username).Find(&user).Error
	if err != nil {
		panic(err)
	}
	return
}

func CreateUser(username string, password string) (user User, err error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return user, err
	}
	user = User{
		Username:  username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = db.DB.Create(&user).Error
	if err != nil {
		panic(err)
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
