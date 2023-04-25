package schedule

import (
	"github.com/gofrs/uuid"
	"time"
)

type Schedule struct {
	ID        int64     `db:"id"`
	DayOfWeek string    `db:"day_of_week"`
	FeedTime  time.Time `db:"feed_time"`
}

type FeedingSchedule struct {
	ID         int64     `db:"id"`
	DeviceID   uuid.UUID `db:"device_id"`
	ScheduleID int       `db:"schedule_id"`
}
