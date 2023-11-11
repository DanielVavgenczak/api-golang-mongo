package entity

import (
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID string `json:"id" bson:"id"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname string `json:"lastname" bson:"lastname"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func NewUser(firstname, lastname, email, password string)*User {
	hashPass, err := hashPassword(password)
	if err != nil {
		panic(err)
	}
	return &User{
		ID: string(uuid.NewString()),
		Firstname: firstname,
		Lastname: lastname,
		Email: email,
		Password: hashPass,
		CreatedAt: time.Now(),
	}
}

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passwordHash), err
}