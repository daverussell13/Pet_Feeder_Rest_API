package feeder

import (
	"context"
	"github.com/daverussell13/Pet_Feeder_Rest_API/internal/connections"
	"strconv"
	"time"
)

type service struct {
	mqtt        *connections.Mqtt
	feedTimeout time.Duration
}

func NewService(mqtt *connections.Mqtt) Service {
	return &service{
		mqtt:        mqtt,
		feedTimeout: time.Duration(3) * time.Second,
	}
}

func (s *service) RealtimeFeed(c context.Context, request *RealtimeFeedRequest) (*RealtimeFeedResponse, error) {
	feedAmount := request.FeedAmount

	mqttClient := s.mqtt.GetClient()
	token := mqttClient.Publish(s.mqtt.GetFeedTopic(), 2, false, strconv.Itoa(feedAmount))
	token.WaitTimeout(s.feedTimeout)

	if token.Error() != nil {
		return nil, token.Error()
	}

	return &RealtimeFeedResponse{
		DeviceID:   request.DeviceID,
		FeedAmount: feedAmount,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
