package tokenValueObject

// TODO:役割的に、valueObjectではなくて、utilsに配置するべきだと思う。
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
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
	cusKey := generateCustomKey(userId, tokenHash)
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

func generateCustomKey(userID string, tokenHash string) string {

	// data := userID + tokenHash
	h := hmac.New(sha256.New, []byte(tokenHash))
	h.Write([]byte(userID))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func GetToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		return "", errors.New("Cookie not found")
	}

	return cookie.Value, nil
}

// ValidateRefreshToken parses and validates the given refresh token
// returns the userId and customkey present in the token payload
func ValidateRefreshToken(tokenString string) (string, string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			fmt.Println("Unexpected signing method in auth token")
			return nil, errors.New("Unexpected signing method in auth token")
		}
		accessTokenKeyPath := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
		verifyBytes, err := ioutil.ReadFile(accessTokenKeyPath)
		if err != nil {
			fmt.Println("unable to read public key")
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			fmt.Println("unable to parse public key")
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		fmt.Println("unable to parse claims")
		return "", "", err
	}

	claims, ok := token.Claims.(*RefreshTokenCustomClaims)
	if !ok || !token.Valid || claims.UserID == "" || claims.KeyType != "refresh" {
		fmt.Println("could not extract claims from token")
		return "", "", errors.New("invalid token: authentication failed")
	}
	return claims.UserID, claims.CustomKey, nil
}

// ValidateAccessToken parses and validates the given access token
// returns the userId present in the token payload
func ValidateAccessToken(tokenString string) (string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			fmt.Println("Unexpected signing method in auth token")
			return nil, errors.New("Unexpected signing method in auth token")
		}

		filePath, err := filepath.Abs(".env")
		if err != nil {
			fmt.Println(err)
		}

		err = godotenv.Load(fmt.Sprintf(filePath))
		if err != nil {
			log.Fatalf("godotenvが使用できません。godotenvをロードしてください。", err)
		}

		accessTokenKeyPath := os.Getenv("ACCESS_TOKEN_PUBLIC_KEY")
		fmt.Println("ACCESS_TOKEN_PUBLIC_KEY" + accessTokenKeyPath)
		verifyBytes, err := ioutil.ReadFile(accessTokenKeyPath)
		if err != nil {
			fmt.Println("unable to read public key")
			return nil, err
		}
		fmt.Println(verifyBytes)
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		fmt.Println(verifyKey)
		if err != nil {
			fmt.Println("unable to parse public key")

			fmt.Println(err)
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		fmt.Println("unable to parse claims")
		return "", err
	}

	claims, ok := token.Claims.(*AccessTokenCustomClaims)
	if !ok || !token.Valid || claims.UserID == "" || claims.KeyType != "access" {
		return "", errors.New("invalid token: authentication failed")
	}
	return claims.UserID, nil
}
