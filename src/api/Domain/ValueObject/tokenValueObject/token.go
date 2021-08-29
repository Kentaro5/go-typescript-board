package tokenValueObject

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// アクセストークンで必要な要素
type AccessTokenCustomClaims struct {
	UserID  string
	KeyType string
	jwt.StandardClaims
}

// リフレッシュトークンで必要な要素
type RefreshTokenCustomClaims struct {
	UserID    string
	CustomKey string
	KeyType   string
	jwt.StandardClaims
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomString generate a string of random characters of given length
func GenerateRandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		idx := rand.Int63() % int64(len(letterBytes))
		sb.WriteByte(letterBytes[idx])
	}
	return sb.String()
}

// アクセストークンの生成
func GenerateAccessToken(userID string) (string, error) {
	tokenType := "access"
	tokenExpiredTime, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRATION"), 10, 64)
	accessTokenKeyPath := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	claims := AccessTokenCustomClaims{
		userID,
		tokenType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(tokenExpiredTime)).Unix(),
			Issuer:    "go-typescript-board.auth.service",
		},
	}

	signBytes, err := ioutil.ReadFile(accessTokenKeyPath)
	if err != nil {
		return "", errors.New("could not generate access token. please try again later")
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("could not generate access token. please try again later")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err := token.SignedString(signKey)

	return accessToken, err
}

func GenerateRefreshToken(userId string, tokenHash string) (string, error) {
	cusKey := GenerateCustomKey(userId, tokenHash)
	tokenType := "refresh"
	accessTokenKeyPath := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")

	claims := RefreshTokenCustomClaims{
		userId,
		cusKey,
		tokenType,
		jwt.StandardClaims{
			Issuer: "go-typescript-board.auth.service",
		},
	}

	signBytes, err := ioutil.ReadFile(accessTokenKeyPath)
	if err != nil {
		return "", errors.New("could not generate refresh token. please try again later")
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", errors.New("could not generate refresh token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err := token.SignedString(signKey)

	return accessToken, err
}

func GenerateCustomKey(userID string, tokenHash string) string {

	// data := userID + tokenHash
	h := hmac.New(sha256.New, []byte(tokenHash))
	h.Write([]byte(userID))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
