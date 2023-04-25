package connections

import (
	"database/sql"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/database"
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

	db, err := database.NewDatabase("pgx", dsn.String())
	if err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func (p *PostgresDB) GetDB() *sql.DB {
	return p.db
}

func (p *PostgresDB) CloseConnection() {
	_ = p.db.Close()
}
