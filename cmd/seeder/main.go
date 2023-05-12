package main

import (
	"database/sql"
	"flag"
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/database/seeders"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func main() {
	var databaseURL string
	flag.StringVar(&databaseURL, "database", "", "database URL")
	flag.Parse()

	if databaseURL == "" {
		log.Fatal("database URL is required")
	}

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Close database gracefully")
	}()

	// Run seeders
	seederList := seeders.SeederList{
		seeders.DeviceSeeder{},
	}

	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	if err = seederList.RunAll(tx); err != nil {
		log.Println("Seeding failed")
		log.Println("Caused by :", err)
		err = tx.Rollback()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Rollback database")
		return
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Seeding success")
}
