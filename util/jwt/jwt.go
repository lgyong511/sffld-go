package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwt相关

const (
	issuer  = "sffld-go"
	subject = "user token"
)

var (
	// jwt密钥
	secret = []byte(time.Now().Format(time.RFC3339Nano))
)

// 自定义claims
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenToken 生成token
func GenToken(username string, timeout int) (string, error) {
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(timeout))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
			Subject:   subject,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
