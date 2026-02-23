package repository

import "database/sql"

type repository struct {
	Sqlite *sql.DB
}

func Newrepository(sqlite *sql.DB) *repository {
	return &repository{Sqlite: sqlite}
}
