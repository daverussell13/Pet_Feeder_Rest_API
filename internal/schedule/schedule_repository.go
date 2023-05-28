package schedule

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/database"
	"github.com/daverussell13/Pet_Feeder_Rest_API/pkg/utils"
	"github.com/gofrs/uuid"
	"time"
)

type repository struct {
	db database.DBTX
}

func NewRepository(db database.DBTX) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) WithTx(ctx context.Context) (Repository, error) {
	if db, ok := r.db.(*sql.DB); ok {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		return NewRepository(tx), nil
	}
	return nil, fmt.Errorf("invalid transaction type")
}

func (r *repository) CommitTx() error {
	if tx, ok := r.db.(*sql.Tx); ok {
		err := tx.Commit()
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("not in a transaction")
}

func (r *repository) RollbackTx() error {
	if tx, ok := r.db.(*sql.Tx); ok {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("not in a transaction")
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

	deviceUUID, err := utils.StringToUUID(deviceId)
	if err != nil {
		return nil, err
	}

	feedingSchedule := FeedingSchedule{
		DeviceID: *deviceUUID,
		Schedule: schedule,
	}

	row := r.db.QueryRowContext(ctx, query, deviceId, schedule.ID)
	err = row.Scan(&feedingSchedule.ID, &feedingSchedule.FeedAmount, &feedingSchedule.CreatedAt)
	if err != nil {
		return &feedingSchedule, err
	}

	return &feedingSchedule, nil
}

func (r *repository) GetDeviceSchedule(ctx context.Context, uuid uuid.UUID) ([]*FeedingSchedule, error) {
	query := `
		SELECT fs.id, fs.device_id, sch.id, sch.day_of_week, sch.feed_time, fs.feed_amount, fs.created_at
		FROM feeding_schedules fs
		JOIN schedules sch ON fs.schedule_id = sch.id
		WHERE fs.device_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, uuid)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var feedingSchedules []*FeedingSchedule
	for rows.Next() {
		feedingSchedule := &FeedingSchedule{}
		schedule := &Schedule{}
		var feedTimeString string
		err = rows.Scan(
			&feedingSchedule.ID,
			&feedingSchedule.DeviceID,
			&schedule.ID,
			&schedule.DayOfWeek,
			&feedTimeString,
			&feedingSchedule.FeedAmount,
			&feedingSchedule.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		schedule.FeedTime, err = time.Parse("15:04:05", feedTimeString)
		if err != nil {
			return nil, err
		}
		feedingSchedule.Schedule = schedule
		feedingSchedules = append(feedingSchedules, feedingSchedule)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return feedingSchedules, nil
}

func (r *repository) GetAllSchedules(ctx context.Context) ([]*Schedule, error) {
	query := "SELECT id, day_of_week, feed_time FROM schedules"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var schedules []*Schedule
	for rows.Next() {
		var feedTimeString string
		schedule := &Schedule{}
		err = rows.Scan(&schedule.ID, &schedule.DayOfWeek, &feedTimeString)
		if err != nil {
			return nil, err
		}
		schedule.FeedTime, err = time.Parse("15:04:05", feedTimeString)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(schedules) == 0 && !rows.Next() {
		return []*Schedule{}, nil
	}

	return schedules, nil
}
