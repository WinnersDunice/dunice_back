package user

import (
	"log"
	"database/sql"
	"crypto/sha256"
	"encoding/hex"
)
type Office struct {
	OfficeID int    `json:"officeid"`
	Address  string `json:"address"`
	Affiliation int  `json:"affiliation"`
}
type User struct {
	ID         int    `json:"id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename"`
	MacAddress string `json:"mac_address"`
}

func HashPassword(password string) (string, error) {
	// Создаем новый хеш
	hash := sha256.New()
	// Записываем пароль в хеш
	hash.Write([]byte(password))
	// Получаем хеш в виде байтового массива
	hashedBytes := hash.Sum(nil)
	// Преобразуем байты в строку в шестнадцатеричном формате
	return hex.EncodeToString(hashedBytes),nil
}

// CreateUser создает нового пользователя
func CreateUser(db *sql.DB, user *User) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	log.Print(user)
	query := `INSERT INTO users (login, password, name, surname, middlename, mac_address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = db.QueryRow(query, user.Login, hashedPassword, user.Name, user.Surname, user.Middlename, user.MacAddress).Scan(&user.ID)
	return err
}
func GetUserByLogin(db *sql.DB, login string) (*User, error) {
	query := `SELECT id, login, name, surname, middlename, mac_address FROM users WHERE login = $1`
	row := db.QueryRow(query, login)
	user := &User{}
	err := row.Scan(&user.ID, &user.Login, &user.Name, &user.Surname, &user.Middlename, &user.MacAddress)
	if err != nil {
		return nil, err
	}
	return user, nil
}
// GetUserByID получает пользователя по ID
func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, login, name, surname, middlename, mac_address FROM users WHERE id = $1`
	row := db.QueryRow(query, id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Login, &user.Name, &user.Surname, &user.Middlename, &user.MacAddress)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser обновляет информацию о пользователе
func UpdateUser(db *sql.DB, id int, updatedUser *User) error {
	query := `UPDATE users SET name = $1, surname = $2, middlename = $3, mac_address = $4 WHERE id = $5`
	_, err := db.Exec(query,updatedUser.Name, updatedUser.Surname, updatedUser.Middlename, updatedUser.MacAddress, id)
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
	query := `UPDATE users SET mac_address = $1 WHERE id = $2`
	_, err := db.Exec(query, newMacAddress, id)
	return err
}
func AuthUser(db *sql.DB, login, password string) (bool, error) {
	query := `SELECT password FROM users WHERE login = $1`
	var hashedPassword string
	err := db.QueryRow(query, login).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}

	//hashedInputPassword, err := HashPassword(password)
	if err != nil {
		return false, err
	}

	return hashedPassword == password, nil
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
	// IsAdmin проверяет, является ли пользователь админом
func IsAdmin(db *sql.DB, userID, officeID int) (bool, error) {
	query := `SELECT isadmin FROM user_offices WHERE userid = $1 AND officeid = $2`
	var isAdmin bool
	err := db.QueryRow(query, userID, officeID).Scan(&isAdmin)
	if err != nil {
		return false, err
	}
	return isAdmin, nil
}

// MakeAdmin делает пользователя админом
func MakeAdmin(db *sql.DB, userID, officeID int) error {
	query := `INSERT INTO user_offices (userid, officeid, isadmin) VALUES ($1, $2, true)`
	_, err := db.Exec(query, userID, officeID)
	return err
}
func GetAllOffices(db *sql.DB) ([]Office, error) {
	query := `SELECT officeid, address, affiliation FROM offices`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offices []Office
	for rows.Next() {
		var office Office
		if err := rows.Scan(&office.OfficeID, &office.Address, &office.Affiliation); err != nil {
			return nil, err
		}
		offices = append(offices, office)
	}
	return offices, nil
}

// GetUsersByOfficeID retrieves all users of a specific office
func GetUsersByOfficeID(db *sql.DB, officeID int) ([]User, error) {
	query := `SELECT u.id, u.login, u.name, u.surname, u.middlename, u.mac_address
              FROM users u
              JOIN user_offices uo ON u.id = uo.userid
              WHERE uo.officeid = $1`
	rows, err := db.Query(query, officeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Login, &user.Name, &user.Surname, &user.Middlename, &user.MacAddress); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetOfficeByUserID(db *sql.DB, userID int) ([]Office, error) {
	query := `SELECT o.officeid, o.address, o.affiliation
              FROM offices o
              JOIN user_offices uo ON o.officeid = uo.officeid
              WHERE uo.userid = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offices []Office
	for rows.Next() {
		var office Office
		if err := rows.Scan(&office.OfficeID, &office.Address, &office.Affiliation); err != nil {
			return nil, err
		}
		offices = append(offices, office)
	}
	return offices, nil
}

