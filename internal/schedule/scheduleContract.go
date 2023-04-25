package schedule

import "context"

type Handler interface {
}

type Service interface {
	AddSchedule()
}

type Repository interface {
	InsertFeedingSchedule(ctx context.Context, feedingSchedule *FeedingSchedule) (*FeedingSchedule, error)
	GetAllSchedule(ctx context.Context) *[]Schedule
}
