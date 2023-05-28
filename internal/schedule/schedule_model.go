package schedule

type ScheduleFeedRequest struct {
	DeviceID   string `json:"device_id" binding:"required,uuid4"`
	DayOfWeek  string `json:"day_of_week" binding:"required,oneof=Sunday Monday Tuesday Wednesday Thursday Friday Saturday"`
	FeedTime   string `json:"feed_time" binding:"required,feedTimeFormat"`
	FeedAmount int    `json:"feed_amount" binding:"required,min=1,max=7"`
}

type ScheduleFeedResponse struct {
	ScheduleID int    `json:"schedule_id" binding:"required"`
	CreatedAt  string `json:"created_at" binding:"required"`
}

type DeviceScheduleListResponse struct {
	Schedule []*FeedingScheduleJson `json:"schedules"`
}

type ScheduleListResponse struct {
	Schedules []*ScheduleJson `json:"schedules"`
}

type ScheduleJson struct {
	ID        int    `json:"id" binding:"required"`
	DayOfWeek string `json:"day_of_week" binding:"required,oneof=Sunday Monday Tuesday Wednesday Thursday Friday Saturday"`
	FeedTime  string `json:"feed_time" binding:"required,feedTimeFormat"`
}

type FeedingScheduleJson struct {
	ID         int    `json:"id" binding:"required"`
	DayOfWeek  string `json:"day_of_week" binding:"required,oneof=Sunday Monday Tuesday Wednesday Thursday Friday Saturday"`
	FeedTime   string `json:"feed_time" binding:"required,feedTimeFormat"`
	FeedAmount int    `json:"feed_amount" binding:"required"`
}

type ScheduleMQTTPayload struct {
	Day    string `json:"day"`
	Hour   int    `json:"hour"`
	Minute int    `json:"minute"`
	Amount int    `json:"amount"`
}
