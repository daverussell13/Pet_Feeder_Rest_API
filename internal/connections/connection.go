package connections

type Connections struct {
	*Mqtt
	*PostgresDB
}

func NewConnection() *Connections {
	mqtt, err := NewMqtt()
	if err != nil {
		panic("Failed to connect to mqtt broker : " + err.Error())
	}

	pg, err := NewPostgresDB()
	if err != nil {
		panic("Failed to connect to postgres database : " + err.Error())
	}

	return &Connections{
		mqtt,
		pg,
	}
}

func (c *Connections) CloseAllConnection() {
	c.Mqtt.CloseConnection()
	c.PostgresDB.CloseConnection()
}
