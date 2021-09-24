package userEntity

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name         string
	Email        string
	PasswordHash string
	TokenHash    string
	SexCode      uint8
	PrefCode     uint32
	CityCode     uint32
	WardCode     uint32
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    *string // may be null
}

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("unable to hash password", "error", err)
		return "", err
	}

	return string(hashedPass), nil
}
