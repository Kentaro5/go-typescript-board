package PasswordUseCase

import (
	"errors"
	"log"

	"api/db"
	"api/infrastructure/userRepositopry"
	"api/utils"
)

func ChangePassword(userId int, oldPassword string, newPassword string) error {
	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
		return err
	}

	// 送られてきたパスワードとIDで、ユーザーが存在するかチェック
	hashedOldPassword, err := utils.HashPassword(oldPassword)
	if err != nil {
		return err
	}

	result, err := userRepositopry.ExistByUserIdAndPassword(connection, userId, hashedOldPassword)
	if err != nil {
		return err
	}

	if !result {
		return errors.New("there is no user by " + utils.IntToString(userId))
	}

	// 新しいパスワードをハッシュ化
	hashedNewPassword, err := utils.HashPassword(newPassword)
	// userRepoで新しいパスワード登録
	err = userRepositopry.UpdatePasswordByUserId(connection, userId, hashedNewPassword)
	if err != nil {
		log.Fatalf("err:", err)
		return err
	}

	return err
}
