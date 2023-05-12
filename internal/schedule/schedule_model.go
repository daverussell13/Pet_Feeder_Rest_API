package schedule

type ScheduledFeedRequest struct {
	DeviceID   string `json:"device_id" binding:"required,uuid4"`
	DayOfWeek  string `json:"day_of_week" binding:"required,oneof=Sunday Monday Tuesday Wednesday Thursday Friday Saturday"`
	FeedTime   string `json:"feed_time" binding:"required,feedTimeFormat"`
	FeedAmount int    `json:"feed_amount" binding:"required,min=1,max=7"`
}

type ScheduledFeedResponse struct {
	ScheduleID string `json:"schedule_id" binding:"required"`
	CreatedAt  string `json:"created_at" binding:"required"`
}

type ListScheduleResponse struct {
	Schedules []*ScheduleJson `json:"schedules"`
}

type ScheduleJson struct {
	ID        int    `db:"id" json:"id" binding:"required"`
	DayOfWeek string `db:"day_of_week" json:"day_of_week" binding:"required,oneof=Sunday Monday Tuesday Wednesday Thursday Friday Saturday"`
	FeedTime  string `db:"feed_time" json:"feed_time" binding:"required,feedTimeFormat"`
}
