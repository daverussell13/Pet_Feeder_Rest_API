package database

import (
	"database/sql"
	"log"
	"net/url"
	"os"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	dsn := url.URL{
		Scheme: os.Getenv("PG_DRIVER"),
		Host:   os.Getenv("PG_HOST"),
		User:   url.UserPassword(os.Getenv("PG_USER"), os.Getenv("PG_PASS")),
		Path:   os.Getenv("PG_DB"),
	}

	queries := dsn.Query()
	queries.Add("sslmode", "disable")

	dsn.RawQuery = queries.Encode()

	db, err := NewDatabase("pgx", dsn.String())
	if err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func (p *PostgresDB) BeginTx() *sql.Tx {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println("Failed to start transaction")
	}
	return tx
}

func (p *PostgresDB) GetDB() *sql.DB {
	return p.db
}

func (p *PostgresDB) CloseConnection() {
	_ = p.db.Close()
}
