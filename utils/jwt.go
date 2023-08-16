package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int64, username string, email string) (string, error) {
	payload := jwt.MapClaims{
		"id":       userId,
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 240).Unix(), // 有效期为24小时
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte("zhengzhong"))
}
