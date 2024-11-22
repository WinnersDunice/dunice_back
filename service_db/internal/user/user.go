package user

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int    `json:"id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename"`
	OfficeID   int    `json:"officeid"`
	MacAddress string `json:"macaddress"`
}

// HashPassword хеширует пароль с использованием bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CreateUser создает нового пользователя
func CreateUser(db *sql.DB, user *User) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (login, password, name, surname, middlename, officeid, macaddress) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err = db.QueryRow(query, user.Login, hashedPassword, user.Name, user.Surname, user.Middlename, user.OfficeID, user.MacAddress).Scan(&user.ID)
	return err
}

// GetUserByID получает пользователя по ID
func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, login, name, surname, middlename, officeid, macaddress FROM users WHERE id = $1`
	row := db.QueryRow(query, id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Login, &user.Name, &user.Surname, &user.Middlename, &user.OfficeID, &user.MacAddress)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser обновляет информацию о пользователе
func UpdateUser(db *sql.DB, id int, updatedUser *User) error {
	query := `UPDATE users SET login = $1, name = $2, surname = $3, middlename = $4, officeid = $5, macaddress = $6 WHERE id = $7`
	_, err := db.Exec(query, updatedUser.Login, updatedUser.Name, updatedUser.Surname, updatedUser.Middlename, updatedUser.OfficeID, updatedUser.MacAddress, id)
	return err
}

// UpdateUserLogin обновляет логин пользователя
func UpdateUserLogin(db *sql.DB, id int, newLogin string) error {
	query := `UPDATE users SET login = $1 WHERE id = $2`
	_, err := db.Exec(query, newLogin, id)
	return err
}

// UpdateUserPassword обновляет пароль пользователя
func UpdateUserPassword(db *sql.DB, id int, newPassword string) error {
	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		return err
	}
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err = db.Exec(query, hashedPassword, id)
	return err
}

// UpdateUserMacAddress обновляет MAC адрес пользователя
func UpdateUserMacAddress(db *sql.DB, id int, newMacAddress string) error {
	query := `UPDATE users SET macaddress = $1 WHERE id = $2`
	_, err := db.Exec(query, newMacAddress, id)
	return err
}

// GetUserLogin получает логин пользователя по ID
func GetUserLogin(db *sql.DB, id int) (string, error) {
	query := `SELECT login FROM users WHERE id = $1`
	var login string
	err := db.QueryRow(query, id).Scan(&login)
	if err != nil {
		return "", err
	}
	return login, nil
}

// GetUserPassword получает хешированный пароль пользователя по ID
func GetUserPassword(db *sql.DB, id int) (string, error) {
	query := `SELECT password FROM users WHERE id = $1`
	var password string
	err := db.QueryRow(query, id).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

// DeleteUser удаляет пользователя по ID
func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}
