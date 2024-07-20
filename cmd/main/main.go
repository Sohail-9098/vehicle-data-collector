package main

import (
	"log"

	"github.com/Sohail-9098/vehicle-data-collector/internal/config"
	"github.com/Sohail-9098/vehicle-data-collector/internal/mqtt"
)

func main() {
	mqttConfig, err := config.NewMQTTConfig()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	client := mqtt.NewClient(mqttConfig.MQTT.Broker, mqttConfig.MQTT.ClientID, mqttConfig.MQTT.Username, mqttConfig.MQTT.Password)
	client.Connect()
	client.Disconnect()
}
