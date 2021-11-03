package TokenUseCase

import (
	"api/db"
	"api/infrastructure/disabledRefreshTokenRepository"
	"api/infrastructure/userRepositopry"
	"api/utils"
	"fmt"
	"log"
)

func RefreshToken(refreshToken string) (string, error) {
	connection, err := db.NewConnection()
	if err != nil {
		log.Fatalf("err:", err)
		return "", err
	}

	// リフレッシュトークンが失効されていないかチェック。
	disabled, _ := disabledRefreshTokenRepository.Exist(connection, refreshToken)
	if disabled {
		return "", err
	}

	// 正しいリフレッシュトークンかチェック。
	claims, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	// 正しいuserIdかチェックする。
	user, err := userRepositopry.FetchByUserId(connection, claims.UserID)
	if err != nil && user.Id != claims.UserID {
		return "", err
	}

	// 使用したリフレッシュトークンを失効
	disabledRefreshTokenRepository.AddDisabledRefreshToken(connection, refreshToken)
	// 新しいリフレッシュトークンを作成
	newRefreshToken, err := utils.GenerateRefreshToken(user.Id, user.TokenHash)
	if err != nil {
		return "", err
	}

	return newRefreshToken, nil
}
