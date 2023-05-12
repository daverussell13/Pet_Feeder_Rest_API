package seeders

import "database/sql"

type Seeder interface {
	Run(tx *sql.Tx) error
}

type SeederList []Seeder

func (sl SeederList) RunAll(tx *sql.Tx) error {
	for _, s := range sl {
		err := s.Run(tx)
		if err != nil {
			return err
		}
	}
	return nil
}
