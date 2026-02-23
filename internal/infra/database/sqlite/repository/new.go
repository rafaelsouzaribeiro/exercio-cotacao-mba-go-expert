package repository

import "database/sql"

type Repository struct {
	Sqlite *sql.DB
}

func NewRepository(sqlite *sql.DB) *Repository {
	return &Repository{Sqlite: sqlite}
}
