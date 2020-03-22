package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"ytu/ginessential/model"
)

var jwtKey = []byte("aSecretCrect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseTocken 生成JWT
func ReleaseTocken(user model.User) (string, error) {
	expirationTime := 7 * 24 * time.Hour
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
			Issuer:    "wzzYtu",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	return token, claims, err
}
