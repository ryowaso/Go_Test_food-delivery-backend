package utils

import "golang.org/x/crypto/bcrypt"

// 加密
func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

// 验证
func CheckPassword(password string, password2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password2), []byte(password))
	return err == nil
}
