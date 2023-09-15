package odontologo

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewRepositoryMySql(db *sql.DB) Repository {
	return &  repository{
		db: db,
	}
}
