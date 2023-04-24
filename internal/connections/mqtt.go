package connections

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

const MqttConnectionTimeout = 3 * time.Second

type Mqtt struct {
	feedTopic string
	client    mqtt.Client
}

func NewMqtt() (*Mqtt, error) {
	server := "tcp://" + os.Getenv("MQTT_BROKER") + ":" + os.Getenv("MQTT_PORT")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(server)
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))

	client := mqtt.NewClient(opts)
	token := client.Connect()

	if !token.WaitTimeout(MqttConnectionTimeout) && token.Error() != nil {
		return nil, token.Error()
	}

	return &Mqtt{
		client:    client,
		feedTopic: "damskuy/petfeeder/feed",
	}, nil
}

func (m *Mqtt) GetClient() mqtt.Client {
	return m.client
}

func (m *Mqtt) CloseConnection() {
	m.client.Disconnect(250)
}

func (m *Mqtt) GetFeedTopic() string {
	return m.feedTopic
}
