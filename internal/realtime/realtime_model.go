package realtime

type FeedRequest struct {
	DeviceID   string `json:"device_id" binding:"required"`
	FeedAmount int    `json:"feed_amount" binding:"required,min=1,max=7"`
}

type FeedResponse struct {
	DeviceID   string `json:"device_id" binding:"required"`
	FeedAmount int    `json:"feed_amount" binding:"required,min=1,max=7"`
	CreatedAt  string `json:"created_at" binding:"required"`
}
