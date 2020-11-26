package db

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Register struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ResetPasword struct {
	ID              string `json:"id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateReset struct {
	Email string `json:"email"`
}

type User struct {
	// ID        string `json:"id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Name      string `json:"Name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func HashPassword(user *Register) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)
}

func CreateHashedPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
