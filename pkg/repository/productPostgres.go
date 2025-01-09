package repository

import (
	"fmt"
	"school21_project1/types"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (p ProductPostgres) Create(product types.Product) (int, error) {
	var uid uuid.NullUUID
	var id int

	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	if len(product.Image) != 0 {
		uid.UUID = uuid.New()
		uid.Valid = true
		query := fmt.Sprintf("INSERT INTO %s (image_id, image) VALUES ($1, $2)", imageTable)
		_, err := p.db.Exec(query, uid.UUID, product.Image)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	query := fmt.Sprintf(`INSERT INTO %s (name,category,price,available_stock, last_update_date, supplier_id, image_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING product_id`, productTable)
	raw := p.db.QueryRow(query, product.Name, product.Category, product.Price, product.AvailableStock, product.LastUpdateDate, product.SupplierID, uid)
	err = raw.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (p ProductPostgres) GetByID(id int) (types.ProductDAO, error) {
	var output types.ProductDAO

	query := fmt.Sprintf(`SELECT p.name, p.category, p.price, p.available_stock, p.last_update_date, s.name, a.country, a.city, a.street, s.phone_number, p.image_id 
	FROM %s p JOIN %s s USING(supplier_id) JOIN %s a USING(adress_id) WHERE p.product_id = $1`, productTable, supplierTable, adressTable)
	err := p.db.Get(&output, query, id)

	return output, err

	// add image later
}

func (p ProductPostgres) GetAll(offset int, limit int) ([]types.ProductDAO, error) {
	var output []types.ProductDAO

	query := fmt.Sprintf(`SELECT p.name, p.category, p.price, p.available_stock, p.last_update_date, s.name, a.country, a.city, a.street, s.phone_number, p.image_id 
	FROM %s p JOIN %s s USING(supplier_id) JOIN %s a USING(adress_id) OFFSET $1 LIMIT $2`, productTable, supplierTable, adressTable)

	err := p.db.Select(&output, query, offset, limit)

	return output, err
}

func (p ProductPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE product_id = $1", productTable)
	_, err := p.db.Exec(query, id)
	return err
}

func (p ProductPostgres) Update(id int, productU types.ProductUpdate) error {
	query := fmt.Sprintf("UPDATE %s SET available_stock = CASE WHEN available_stock - $1 >= 0 THEN available_stock - $2 ELSE available_stock END WHERE product_id = $3", productTable)

	_, err := p.db.Exec(query, productU.AvailableStockU, productU.AvailableStockU, id)

	return err
}
