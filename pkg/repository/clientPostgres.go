package repository

import (
	"fmt"
	"strings"

	"school21_project1/types"

	"github.com/jmoiron/sqlx"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (c ClientPostgres) Create(client types.CreateClient) (int, error) {
	var id int

	tx, err := c.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(`INSERT INTO %s (country, city, street) VALUES ($1, $2, $3) RETURNING adress_id`, adressTable)
	raw := c.db.QueryRow(query, client.Country, client.City, client.Street)
	if err := raw.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(`INSERT INTO %s (name, surname, birthday, gender, registration_date, adress_id)
    VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5) RETURNING client_id`, clientTable)
	raw = c.db.QueryRow(query, client.Name, client.Surname, client.Birthday, client.Gender, id)
	if err := raw.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (c ClientPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE client_id=$1", clientTable)
	_, err := c.db.Exec(query, id)
	return err
}

func (c ClientPostgres) Find(name string, surname string) ([]types.ClientDTO, error) {
	var output []types.ClientDTO

	query := fmt.Sprintf("SELECT client_id, name, surname, birthday, gender, registration_date, adress_id FROM %s WHERE name=$1 AND surname=$2", clientTable)
	err := c.db.Select(&output, query, name, surname)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c ClientPostgres) GetAll(limit int, offset int) ([]types.ClientDTO, error) {
	var output []types.ClientDTO

	query := fmt.Sprintf("SELECT client_id, name, surname, birthday, gender, registration_date, adress_id FROM %s OFFSET $1 LIMIT $2", clientTable)
	err := c.db.Select(&output, query, offset, limit)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (c ClientPostgres) Update(id int, adress types.Adress) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if adress.Country != nil {
		setValues = append(setValues, fmt.Sprintf("country=$%d", argID))
		args = append(args, *adress.Country)
		argID++
	}

	if adress.City != nil {
		setValues = append(setValues, fmt.Sprintf("city=$%d", argID))
		args = append(args, *adress.City)
		argID++
	}

	if adress.Street != nil {
		setValues = append(setValues, fmt.Sprintf("street=$%d", argID))
		args = append(args, *adress.Street)
		argID++
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf(`UPDATE %s ad SET %s FROM %s cl WHERE ad.adress_id = cl.adress_id AND cl.client_id = $%d`, adressTable, setQuery, clientTable, argID)

	args = append(args, id)
	_, err := c.db.Exec(query, args...)

	return err
}
