package repository

import (
	"fmt"
	"school21_project1/types"

	"github.com/jmoiron/sqlx"
)

type SupplierPostgres struct {
	db *sqlx.DB
}

func NewSupplierPostgres(db *sqlx.DB) *SupplierPostgres {
	return &SupplierPostgres{db: db}
}

func (s SupplierPostgres) Create(supplier types.CreateSupplier) (int, error) {
	var id int
	tx, err := s.db.Begin()
	if err != nil {
		return 0, nil
	}

	query := fmt.Sprintf("INSERT INTO %s (country, city, street) VALUES ($1, $2, $3) RETURNING adress_id", adressTable)
	raw := s.db.QueryRow(query, supplier.Country, supplier.City, supplier.Street)
	if err := raw.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}

	query = fmt.Sprintf("INSERT INTO %s (supplier_name, adress_id, supplier_phone_number) VALUES ($1, $2, $3) RETURNING supplier_id", supplierTable)
	raw = s.db.QueryRow(query, supplier.Name, id, supplier.PhoneNumber)
	if err := raw.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}

	return id, tx.Commit()
}

func (s SupplierPostgres) Update(id int, adress types.Adress) error {
	query := fmt.Sprintf("UPDATE %s ad USING %s sp SET ad.country = $1, ad.city = $2, ad.street = $3 WHERE ad.adress_id = sp.adress_id AND sp.supplier_id = $4", adressTable, supplierTable)
	_, err := s.db.Exec(query, adress.Country, adress.City, adress.Street, id)

	return err
}

func (s SupplierPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE supplier_id = $1", supplierTable)
	_, err := s.db.Exec(query, id)

	return err
}

func (s SupplierPostgres) GetAll() ([]types.SupplierDTO, error) {
	var supplier []types.SupplierDTO

	query := fmt.Sprintf("SELECT sp.supplier_name, ad.country, ad.city, ad.street, sp.supplier_phone_number FROM %s sp JOIN %s ad USING(adress_id)", supplierTable, adressTable)
	err := s.db.Select(&supplier, query)

	return supplier, err
}

func (s SupplierPostgres) GetByID(id int) (types.SupplierDTO, error) {
	var supplier types.SupplierDTO

	query := fmt.Sprintf("SELECT sp.supplier_name, ad.country, ad.city, ad.street, sp.supplier_phone_number FROM %s sp JOIN %s ad USING(adress_id) WHERE sp.supplier_id = $1", supplierTable, adressTable)

	err := s.db.Get(&supplier, query, id)

	return supplier, err
}
