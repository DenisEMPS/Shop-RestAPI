package repository

import (
	"fmt"
	"school21_project1/types"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ImagePostgres struct {
	db *sqlx.DB
}

func NewImagePostgres(db *sqlx.DB) *ImagePostgres {
	return &ImagePostgres{db: db}
}

func (i ImagePostgres) Create(image types.CreateImageProduct) (string, error) {
	uuid := uuid.New()

	tx, err := i.db.Begin()
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("INSERT INTO %s (image_id, image) VALUES ($1, $2)", imageTable)
	_, err = i.db.Exec(query, uuid, image.Image)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	query = fmt.Sprintf("UPDATE %s SET image_id=$1 WHERE product_id = $2", productTable)
	_, err = i.db.Exec(query, uuid, image.ProductID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return uuid.String(), tx.Commit()
}

func (i ImagePostgres) GetByID(id string) (types.Image, error) {
	var imageByID types.Image

	query := "SELECT image FROM " + imageTable + " WHERE image_id = $1"

	err := i.db.Get(&imageByID.Image, query, id)
	if err != nil {
		return imageByID, fmt.Errorf("could not get image by id: %w", err)
	}

	if len(imageByID.Image) == 0 {
		return imageByID, fmt.Errorf("image data is empty for id: %s", id)
	}

	return imageByID, nil
}

func (i ImagePostgres) Update(uuid string, image types.Image) error {
	query := fmt.Sprintf("UPDATE %s SET image=$1 WHERE image_id = $2", imageTable)
	_, err := i.db.Exec(query, image.Image, uuid)

	return err
}

func (i ImagePostgres) Delete(uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE image_id = $1", imageTable)
	_, err := i.db.Exec(query, uuid)

	return err
}

func (i ImagePostgres) GetByProductID(id int) (types.Image, error) {
	var image types.Image

	query := fmt.Sprintf("SELECT image FROM %s JOIN %s USING(image_id) WHERE product_id = $1", imageTable, productTable)
	err := i.db.Get(&image.Image, query, id)

	if len(image.Image) == 0 {
		return image, fmt.Errorf("there is no image for product id: %v", id)
	}

	return image, err
}
