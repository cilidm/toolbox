package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("20060102$%#my_pack&*^150405")

type Claims struct {
	UserID    int    `json:"user_id"`
	LoginName string `json:"login_name"`
	jwt.StandardClaims
}

func GenerateToken(uid int, loginName string) (token string, err error) {
	now := time.Now()
	expireTime := now.Add(3 * time.Hour)
	claims := Claims{
		uid,
		loginName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "oss_manage",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
