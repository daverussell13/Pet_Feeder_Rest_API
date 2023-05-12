package schedule

import (
	"context"
	"database/sql"
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/mqtt"
	"github.com/daverussell13/Pet_Feeder_Rest_API/pkg/utils"
	"strconv"
	"time"
)

type service struct {
	scheduleRepository Repository
	addScheduleTimeout time.Duration
	mqtt               *mqtt.Mqtt
}

func NewService(scheduleRepository Repository, mqtt *mqtt.Mqtt) Service {
	return &service{
		mqtt:               mqtt,
		scheduleRepository: scheduleRepository,
		addScheduleTimeout: time.Duration(5) * time.Second,
	}
}

func (s *service) AddSchedule(c context.Context, req *ScheduledFeedRequest) (*ScheduledFeedResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.addScheduleTimeout)
	defer cancel()

	day := req.DayOfWeek
	feedTime := utils.StringToTime(req.FeedTime)

	scheduleRepositoryTx, err := s.scheduleRepository.WithTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = scheduleRepositoryTx.RollbackTx()
	}()

	schedule, err := s.scheduleRepository.GetScheduleByDayAndTime(ctx, day, feedTime)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		schedule, err = scheduleRepositoryTx.InsertSchedule(ctx, schedule)
		if err != nil {
			return nil, err
		}
	}

	feedingSchedule, err := s.scheduleRepository.GetSameScheduleOnDevice(ctx, req.DeviceID, schedule)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		feedingSchedule.FeedAmount = int8(req.FeedAmount)
		feedingSchedule, err = scheduleRepositoryTx.InsertFeedingSchedule(ctx, feedingSchedule)
		if err != nil {
			return nil, err
		}
	}

	err = scheduleRepositoryTx.CommitTx()
	if err != nil {
		return nil, err
	}

	return &ScheduledFeedResponse{
		ScheduleID: strconv.Itoa(int(feedingSchedule.ID)),
		CreatedAt:  feedingSchedule.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *service) ShowAllSchedules(ctx context.Context) (*ListScheduleResponse, error) {
	c, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Second)
	defer cancel()

	schedules, err := s.scheduleRepository.GetAllSchedules(c)
	if err != nil {
		return nil, err
	}

	var schedulesJson []*ScheduleJson
	for _, schedule := range schedules {
		scheduleJson := &ScheduleJson{
			ID:        int(schedule.ID),
			DayOfWeek: schedule.DayOfWeek,
			FeedTime:  schedule.FeedTime.Format("15:04"),
		}
		schedulesJson = append(schedulesJson, scheduleJson)
	}

	return &ListScheduleResponse{
		Schedules: schedulesJson,
	}, err
}
