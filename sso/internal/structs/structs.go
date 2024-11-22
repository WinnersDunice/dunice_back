package structs

import (
	"crypto/sha256"
	"encoding/hex"
)

var IPDB string = "http://185.23.236.113:8003/database"

type User struct {
	Email    string `json:"email"`
	UserID   int    `json:"id"`
	Password string `json:"password"`
	Role     int    `json:"privilege_level"`
}

func HashPassword(password string) string {
	// Создаем новый хеш
	hash := sha256.New()
	// Записываем пароль в хеш
	hash.Write([]byte(password))
	// Получаем хеш в виде байтового массива
	hashedBytes := hash.Sum(nil)
	// Преобразуем байты в строку в шестнадцатеричном формате
	return hex.EncodeToString(hashedBytes)
}
