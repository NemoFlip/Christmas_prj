package database

import "database/sql"

type ProductsStorage struct {
	db *sql.DB
}

func NewProductsStorage(db *sql.DB) *ProductsStorage {
	return &ProductsStorage{db: db}
}
