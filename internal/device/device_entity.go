package device

import (
	"github.com/gofrs/uuid"
)

type Device struct {
	ID         uuid.UUID `db:"id"`
	DeviceName string    `db:"device_name"`
	DeviceType string    `db:"device_type"`
}
