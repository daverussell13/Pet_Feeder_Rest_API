package seeders

import "database/sql"

type Seeder interface {
	Run(db *sql.DB) error
}

type SeederList []Seeder

func (sl SeederList) RunAll(db *sql.DB) error {
	for _, s := range sl {
		err := s.Run(db)
		if err != nil {
			return err
		}
	}
	return nil
}
