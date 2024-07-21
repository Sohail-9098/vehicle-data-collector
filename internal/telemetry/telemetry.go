package telemetry

import (
	"log"

	"github.com/Sohail-9098/vehicle-data-collector/internal/config"
	"github.com/Sohail-9098/vehicle-data-collector/internal/mqtt"
	"github.com/Sohail-9098/vehicle-data-collector/internal/protobufs/vehicle"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
)

type Telemetry struct {
	Data []*vehicle.Telemetry
}

func NewTelemetryData() *Telemetry {
	return &Telemetry{}
}

func (t *Telemetry) FetchTelemetryData() {
	mqttConfig, err := config.NewMQTTConfig()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	client := mqtt.NewClient(mqttConfig.MQTT.Broker, mqttConfig.MQTT.ClientID, mqttConfig.MQTT.Username, mqttConfig.MQTT.Password)
	client.Connect()
	defer client.Disconnect()

	if err := client.Subscribe("vehicles/#", 0, t.messageHandler); err != nil {
		log.Fatalf("failed to subscribe: %v", err)
	}
	select {}
}

func (t *Telemetry) messageHandler(c MQTT.Client, m MQTT.Message) {
	telemetry := vehicle.Telemetry{}
	if err := proto.Unmarshal(m.Payload(), &telemetry); err != nil {
		log.Printf("failed to read telemetry data: %v\n", err)
	}
	t.Data = append(t.Data, &telemetry)
}
