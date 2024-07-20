package mqtt

import (
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	client MQTT.Client
}

func NewClient(broker, clientID, username, password string) *Client {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	mqttClient := MQTT.NewClient(opts)
	return &Client{client: mqttClient}
}

func (c *Client) Connect() error {
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	log.Println("MQTT client connected")
	return nil
}

func (c *Client) Disconnect() {
	c.client.Disconnect(250)
	log.Println("MQTT client disconnected")
}
