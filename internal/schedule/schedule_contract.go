package schedule

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler interface {
	ScheduledFeed(ctx *gin.Context)
	ListSchedule(ctx *gin.Context) // TODO: remove me
}

type Service interface {
	AddSchedule(ctx context.Context, request *ScheduledFeedRequest) (*ScheduledFeedResponse, error)
	ShowAllSchedules(ctx context.Context) (*ListScheduleResponse, error) // TODO: remove me
}

type Repository interface {
	InsertSchedule(ctx context.Context, s *Schedule) (*Schedule, error)
	InsertFeedingSchedule(ctx context.Context, s *FeedingSchedule) (*FeedingSchedule, error)
	GetScheduleByDayAndTime(ctx context.Context, day string, time time.Time) (*Schedule, error)
	GetSameScheduleOnDevice(ctx context.Context, deviceId string, schedule *Schedule) (*FeedingSchedule, error)
	GetAllSchedules(ctx context.Context) ([]*Schedule, error)
	WithTx(ctx context.Context) (Repository, error)
	CommitTx() error
	RollbackTx() error
}
