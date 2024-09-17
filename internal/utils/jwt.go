package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("FanOne")

type Claims struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.RegisteredClaims
}

func GenerateToken(id uint64, username string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:        id,
		Username:  username,
		Authority: authority,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "mall",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
