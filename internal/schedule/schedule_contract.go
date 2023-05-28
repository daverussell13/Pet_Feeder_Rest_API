package schedule

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"time"
)

type Handler interface {
	ScheduleFeed(ctx *gin.Context)
	DeviceScheduleList(ctx *gin.Context)
	ScheduleList(ctx *gin.Context) // TODO: remove me
}

type Service interface {
	AddSchedule(ctx context.Context, request *ScheduleFeedRequest) (*ScheduleFeedResponse, error)
	ShowDeviceSchedules(ctx context.Context, deviceId uuid.UUID) (*DeviceScheduleListResponse, error)
	ShowAllSchedules(ctx context.Context) (*ScheduleListResponse, error) // TODO: remove me
}

type Repository interface {
	InsertSchedule(ctx context.Context, s *Schedule) (*Schedule, error)
	InsertFeedingSchedule(ctx context.Context, s *FeedingSchedule) (*FeedingSchedule, error)
	GetScheduleByDayAndTime(ctx context.Context, day string, time time.Time) (*Schedule, error)
	GetSameScheduleOnDevice(ctx context.Context, deviceId string, schedule *Schedule) (*FeedingSchedule, error)
	GetDeviceSchedule(ctx context.Context, uuid uuid.UUID) ([]*FeedingSchedule, error)
	GetAllSchedules(ctx context.Context) ([]*Schedule, error)
	WithTx(ctx context.Context) (Repository, error)
	CommitTx() error
	RollbackTx() error
}
