package realtime

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	RealtimeFeed(ctx *gin.Context)
}

type Service interface {
	RealtimeFeed(ctx context.Context, request *FeedRequest) (*FeedResponse, error)
}
