package seeders

import (
	"database/sql"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/device"
	"github.com/gofrs/uuid"
)

type DeviceSeeder struct{}

func seed() []device.Device {
	return []device.Device{
		{
			ID:         uuid.FromStringOrNil("592f5ec4-c3dd-4d5c-93fe-4ee3db513ad7"),
			DeviceName: "PetFeeder1",
			DeviceType: "Pet Feeder",
		},
		{
			ID:         uuid.FromStringOrNil("a02a7dbc-0b28-4f63-80fa-646f9cf1bff1"),
			DeviceName: "PetFeeder2",
			DeviceType: "Pet Feeder",
		},
	}
}

func (s DeviceSeeder) Run(tx *sql.Tx) error {
	devices := seed()
	query := "INSERT INTO devices (id, device_name, device_type) VALUES ($1, $2, $3)"
	for _, d := range devices {
		_, err := tx.Exec(query, d.ID, d.DeviceName, d.DeviceType)
		if err != nil {
			return err
		}
	}
	return nil
}
