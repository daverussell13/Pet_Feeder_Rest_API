package schedule

import (
	"context"
	"database/sql"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/utils"
	"strconv"
	"time"
)

type service struct {
	scheduleRepository Repository
	addScheduleTimeout time.Duration
}

func NewService(scheduleRepository Repository) Service {
	return &service{
		scheduleRepository: scheduleRepository,
		addScheduleTimeout: time.Duration(5) * time.Second,
	}
}

func (s *service) AddSchedule(c context.Context, req *ScheduledFeedRequest) (*ScheduledFeedResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.addScheduleTimeout)
	defer cancel()

	day := req.DayOfWeek
	feedTime := utils.StringToTime(req.FeedTime)

	schedule, err := s.scheduleRepository.GetScheduleByDayAndTime(ctx, day, feedTime)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		schedule, err = s.scheduleRepository.InsertSchedule(ctx, schedule)
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
		feedingSchedule, err = s.scheduleRepository.InsertFeedingSchedule(ctx, feedingSchedule)
		if err != nil {
			return nil, err
		}
	}

	return &ScheduledFeedResponse{
		ScheduleID: strconv.Itoa(int(feedingSchedule.ID)),
		CreatedAt:  feedingSchedule.CreatedAt.String(),
	}, nil
}
