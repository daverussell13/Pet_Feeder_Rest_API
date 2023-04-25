package database

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewDatabase(driverName string, dataSource string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
