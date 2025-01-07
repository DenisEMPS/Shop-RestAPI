package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	clientTable   = "client"
	productTable  = "product"
	supplierTable = "supplier"
	imageTable    = "images"
	countryTable  = "country"
	adressTable   = "adress"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBname   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
