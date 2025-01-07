package repository

import (
	"fmt"
	"school21_project1/types"

	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (p ProductPostgres) Create(product types.ProductDAO) (int, error) {
	var id int

	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (image) VALUES ($1) RETURNING image_id", imageTable)
	raw := p.db.QueryRow(query, product.Image)
	if err := raw.Scan(&id); err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (country, city, street) VALUES ($1, $2, $3) RETURNING adress_id", adressTable)
	raw := p.db.QueryRow(query, pr)

}
