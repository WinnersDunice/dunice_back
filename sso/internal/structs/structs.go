package structs

import (
	"crypto/sha256"
	"encoding/hex"
)

var IPDB string = "http://195.80.238.9:8003/database"

type User struct {
	Login    string `json:"login"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename"`
	MacAddress string `json:"mac_address"`
	OfficeID int    `json:"officeid"`
	IsAdmin  bool   `json:"isadmin"`
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
