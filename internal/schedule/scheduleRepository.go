package schedule

import (
	"context"
	"database/sql"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/utils"
	"time"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertSchedule(ctx context.Context, s *Schedule) (*Schedule, error) {
	query := "INSERT INTO schedules(day_of_week, feed_time) VALUES ($1, $2) RETURNING id"

	var lastInsertedId int64
	row := r.db.QueryRowContext(ctx, query, s.DayOfWeek, s.FeedTime)

	err := row.Scan(&lastInsertedId)
	if err != nil {
		return nil, err
	}

	s.ID = lastInsertedId
	return s, nil
}

func (r *repository) InsertFeedingSchedule(ctx context.Context, s *FeedingSchedule) (*FeedingSchedule, error) {
	query := "INSERT INTO feeding_schedules(device_id, schedule_id, feed_amount) VALUES ($1, $2, $3) RETURNING id, created_at"

	var lastInsertedId int64
	var createdAt time.Time

	row := r.db.QueryRowContext(ctx, query, s.DeviceID, s.Schedule.ID, s.FeedAmount)

	err := row.Scan(&lastInsertedId, &createdAt)
	if err != nil {
		return nil, err
	}

	s.ID = lastInsertedId
	s.CreatedAt = createdAt
	return s, nil
}

func (r *repository) GetScheduleByDayAndTime(ctx context.Context, day string, time time.Time) (*Schedule, error) {
	query := "SELECT id FROM schedules WHERE day_of_week = $1 AND feed_time = $2"

	schedule := Schedule{
		DayOfWeek: day,
		FeedTime:  time,
	}

	row := r.db.QueryRowContext(ctx, query, day, time)
	err := row.Scan(&schedule.ID)
	if err != nil {
		return &schedule, err
	}

	return &schedule, nil
}

func (r *repository) GetSameScheduleOnDevice(ctx context.Context, deviceId string, schedule *Schedule) (*FeedingSchedule, error) {
	query := `
		SELECT id, feed_amount, created_at
		FROM feeding_schedules
		WHERE device_id = $1 AND schedule_id = $2
	`

	feedingSchedule := FeedingSchedule{
		DeviceID: utils.StringToUUID(deviceId),
		Schedule: schedule,
	}

	row := r.db.QueryRowContext(ctx, query, deviceId, schedule.ID)
	err := row.Scan(&feedingSchedule.ID, &feedingSchedule.FeedAmount, &feedingSchedule.CreatedAt)
	if err != nil {
		return &feedingSchedule, err
	}

	return &feedingSchedule, nil
}
