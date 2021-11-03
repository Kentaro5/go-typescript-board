package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("unable to hash password", "error", err)
		return "", err
	}

	return string(hashedPass), nil
}

func CheckPassword(password string, requestPassword string) error {
	// テキストのパスワードと、ハッシュ化されたパスワードを比較する
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(requestPassword))
	if err != nil {
		return err
	}

	return err
}
