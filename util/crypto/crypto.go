package crypto

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// 密码加密相关

// GenSalt 生成随机盐
func GenSalt() (string, error) {
	b := make([]byte, bcrypt.MaxCost)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// ComparePassword 比较密码
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenPassword 生成密码
func GenPassword(password, salt string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.MaxCost)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
