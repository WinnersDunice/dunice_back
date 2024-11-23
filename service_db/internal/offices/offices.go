package office

import (
	"database/sql"
)

type Office struct {
	OfficeID int    `json:"officeid"`
	Address  string `json:"address"`
	Affiliation int  `json:"affiliation"`
}

// CreateOffice создает новый офис
func CreateOffice(db *sql.DB, office *Office) error {
	query := `INSERT INTO offices (address) VALUES ($1) RETURNING officeid`
	err := db.QueryRow(query, office.Address).Scan(&office.OfficeID)
	return err
}

// GetOfficeByID получает офис по ID
func GetOfficeByID(db *sql.DB, id int) (*Office, error) {
	query := `SELECT officeid, address FROM offices WHERE officeid = $1`
	row := db.QueryRow(query, id)
	office := &Office{}
	err := row.Scan(&office.OfficeID, &office.Address)
	if err != nil {
		return nil, err
	}
	return office, nil
}

// UpdateOffice обновляет информацию об офисе
func UpdateOffice(db *sql.DB, id int, updatedOffice *Office) error {
	query := `UPDATE offices SET address = $1 WHERE officeid = $2`
	_, err := db.Exec(query, updatedOffice.Address, id)
	return err
}

// DeleteOffice удаляет офис по ID
func DeleteOffice(db *sql.DB, id int) error {
	query := `DELETE FROM offices WHERE officeid = $1`
	_, err := db.Exec(query, id)
	return err
}
