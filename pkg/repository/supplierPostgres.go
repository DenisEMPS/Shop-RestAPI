package repository

import (
	"fmt"
	"school21_project1/types"
	"strings"

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
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (country, city, street) VALUES ($1, $2, $3) RETURNING adress_id", adressTable)
	raw := s.db.QueryRow(query, supplier.Country, supplier.City, supplier.Street)
	if err := raw.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (supplier_name, adress_id, supplier_phone_number) VALUES ($1, $2, $3) RETURNING supplier_id", supplierTable)
	raw = s.db.QueryRow(query, supplier.Name, id, supplier.PhoneNumber)
	if err := raw.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (s SupplierPostgres) Update(id int, adress types.AdressDTO) error {
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

	values := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s ad SET %s FROM %s sp WHERE ad.adress_id = sp.adress_id AND sp.supplier_id = $%d", adressTable, values, supplierTable, argID)

	args = append(args, id)
	_, err := s.db.Exec(query, args...)

	return err
}

func (s SupplierPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE supplier_id = $1", supplierTable)
	_, err := s.db.Exec(query, id)

	return err
}

func (s SupplierPostgres) GetAll() ([]types.SupplierDAO, error) {
	var supplier []types.SupplierDAO

	query := fmt.Sprintf("SELECT sp.supplier_name, ad.country, ad.city, ad.street, sp.supplier_phone_number FROM %s sp JOIN %s ad USING(adress_id)", supplierTable, adressTable)
	err := s.db.Select(&supplier, query)

	return supplier, err
}

func (s SupplierPostgres) GetByID(id int) (types.SupplierDAO, error) {
	var supplier types.SupplierDAO

	query := fmt.Sprintf("SELECT sp.supplier_name, ad.country, ad.city, ad.street, sp.supplier_phone_number FROM %s sp JOIN %s ad USING(adress_id) WHERE sp.supplier_id = $1", supplierTable, adressTable)

	err := s.db.Get(&supplier, query, id)

	return supplier, err
}
