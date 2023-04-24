package feeder

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	RealtimeFeed(ctx *gin.Context)
}

type Service interface {
	RealtimeFeed(ctx context.Context, request *RealtimeFeedRequest) (*RealtimeFeedResponse, error)
}
