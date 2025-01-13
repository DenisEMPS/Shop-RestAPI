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
		_, err := p.db.Exec(query, uid, product.Image)
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

func (p ProductPostgres) GetByID(id int) (types.ProductDAO, types.Image, error) {
	var outputProduct types.ProductDAO
	var outputImage types.Image

	query := fmt.Sprintf(`SELECT p.name, p.category, p.price, p.available_stock, p.last_update_date, supplier_name, a.country, a.city, a.street, supplier_phone_number, i.image 
	FROM %s p JOIN %s s USING(supplier_id) JOIN %s a USING(adress_id) JOIN %s i USING(image_id) WHERE p.product_id = $1`, productTable, supplierTable, adressTable, imageTable)
	raw, err := p.db.Query(query, id)
	if err != nil {
		return outputProduct, outputImage, err
	}

	for raw.Next() {
		err = raw.Scan(&outputProduct.Name, &outputProduct.Category, &outputProduct.Price, &outputProduct.AvailableStock, &outputProduct.LastUpdateDate, &outputProduct.SupplierName, &outputProduct.SupplierAdressCountry, &outputProduct.SupplierAdressCity, &outputProduct.SupplierAdressStreet, &outputProduct.SupplierPhoneNumber, &outputImage.Image)
		if err != nil {
			return outputProduct, outputImage, err
		}
	}

	if outputProduct.Name == "" {
		return outputProduct, outputImage, fmt.Errorf("product not found")
	}

	return outputProduct, outputImage, err
}

func (p ProductPostgres) GetAll(offset int, limit int) ([]types.ProductDAO, []types.Image, error) {
	var outputProducts []types.ProductDAO
	var outputImages []types.Image

	query := fmt.Sprintf(`SELECT p.name, p.category, p.price, p.available_stock, p.last_update_date, supplier_name, a.country, a.city, a.street, supplier_phone_number, image
	FROM %s p JOIN %s s USING(supplier_id) JOIN %s a USING(adress_id) JOIN %s USING(image_id)`, productTable, supplierTable, adressTable, imageTable)

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}

	if offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", offset)
	}

	raws, err := p.db.Query(query)
	if err != nil {
		return nil, nil, err
	}

	for raws.Next() {
		var outputP types.ProductDAO
		var outputI types.Image
		err := raws.Scan(&outputP.Name, &outputP.Category, &outputP.Price, &outputP.AvailableStock, &outputP.LastUpdateDate, &outputP.SupplierName, &outputP.SupplierAdressCountry, &outputP.SupplierAdressCity, &outputP.SupplierAdressStreet, &outputP.SupplierPhoneNumber, &outputI.Image)
		if err != nil {
			return nil, nil, err
		}
		outputProducts = append(outputProducts, outputP)
		outputImages = append(outputImages, outputI)
	}

	if outputProducts == nil {
		return nil, nil, fmt.Errorf("no products was find")
	}

	return outputProducts, outputImages, nil
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
