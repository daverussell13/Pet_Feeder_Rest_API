package realtime

import (
	"context"
	"errors"
	"github.com/daverussell13/Pet_Feeder_Rest_API/infrastructures/mqtt"
	"github.com/daverussell13/Pet_Feeder_Rest_API/pkg/server"
	mqttLib "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strconv"
	"time"
)

type service struct {
	mqtt        *mqtt.Mqtt
	feedTimeout time.Duration
}

func NewService(mqtt *mqtt.Mqtt) Service {
	return &service{
		mqtt:        mqtt,
		feedTimeout: time.Duration(3) * time.Second,
	}
}

func (s *service) RealtimeFeed(c context.Context, request *FeedRequest) (*FeedResponse, error) {
	feedAmount := request.FeedAmount
	mqttClient := s.mqtt.GetClient()

	topic := s.mqtt.GetTopic().FeedTopic + "/" + request.DeviceID
	ackTopic := topic + "/acknowledge"

	subChan := make(chan string)
	subToken := mqttClient.Subscribe(ackTopic, 0, func(client mqttLib.Client, msg mqttLib.Message) {
		subChan <- string(msg.Payload())
	})
	defer mqttClient.Unsubscribe(ackTopic)

	if !subToken.WaitTimeout(time.Duration(2) * time.Second) {
		return nil, errors.New(ackTopic + " Subscribe timeout")
	}
	if err := subToken.Error(); err != nil {
		return nil, err
	}

	token := mqttClient.Publish(topic, 2, false, strconv.Itoa(feedAmount))
	if !token.WaitTimeout(time.Duration(2) * time.Second) {
		return nil, errors.New(topic + " Publish timeout")
	}
	if err := token.Error(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c, s.feedTimeout)
	defer cancel()

	select {
	case message := <-subChan:
		log.Println(message)
		return &FeedResponse{
			DeviceID:   request.DeviceID,
			FeedAmount: feedAmount,
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		}, nil
	case <-ctx.Done():
		return nil, errors.New(server.DeviceUnresponsive)
	}
}
