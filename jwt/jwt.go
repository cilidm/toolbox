package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ClaimsOut struct {
	UserId    int    `json:"user_id"`
	LoginName string `json:"login_name"`
}

type Claims struct {
	ClaimsOut
	jwtSecret []byte
	issuer    string
	subject   string
	jwt.StandardClaims
}

func NewClaims(userId int, loginName string) *Claims {
	return &Claims{ClaimsOut: ClaimsOut{UserId: userId, LoginName: loginName},
		jwtSecret: []byte("https://www.github.com/cilidm/toolbox"), issuer: "127.0.0.1", subject: "user-token"}
}

func (c *Claims) SetJwtSecret(secret string) *Claims {
	c.jwtSecret = []byte(secret)
	return c
}

func (c *Claims) SetSubject(subject string) *Claims {
	c.subject = subject
	return c
}

func (c *Claims) SetIssuer(issuer string) *Claims {
	c.issuer = issuer
	return c
}

func (c *Claims) Generate() (token string, err error) {
	expireTime := time.Now().Add(24 * time.Hour)
	c.ExpiresAt = expireTime.Unix()
	c.IssuedAt = time.Now().Unix()
	c.Issuer = c.issuer
	c.Subject = c.subject
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokenClaims.SignedString(c.jwtSecret)
}

func (c *Claims) Parse(token string) (*ClaimsOut, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return c.jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return &claims.ClaimsOut, nil
		}
	}
	return nil, err
}
