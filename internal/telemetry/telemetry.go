package telemetry

import (
	"context"
	"log"
	"sync"

	"github.com/Sohail-9098/vehicle-data-collector/internal/config"
	"github.com/Sohail-9098/vehicle-data-collector/internal/mqtt"
	"github.com/Sohail-9098/vehicle-data-collector/internal/protobufs/vehicle"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type Telemetry struct {
	dataCh chan *vehicle.Telemetry
}

func NewTelemetry() *Telemetry {
	return &Telemetry{
		dataCh: make(chan *vehicle.Telemetry, 100),
	}
}

func (t *Telemetry) FetchAndProcessTelemetryData() {
	mqttConfig, err := config.New()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	client := mqtt.NewClient(mqttConfig.MQTT.Broker, mqttConfig.MQTT.ClientID, mqttConfig.MQTT.Username, mqttConfig.MQTT.Password)
	client.Connect()
	defer client.Disconnect()

	if err := client.Subscribe("vehicles/#", 0, t.messageHandler); err != nil {
		log.Fatalf("failed to subscribe: %v", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go t.processTelemetryData(&wg)
	wg.Wait()
}

func (t *Telemetry) messageHandler(c MQTT.Client, m MQTT.Message) {
	go func() {
		telemetry := vehicle.Telemetry{}
		if err := proto.Unmarshal(m.Payload(), &telemetry); err != nil {
			log.Printf("failed to read telemetry data: %v\n", err)
		}
		t.dataCh <- &telemetry
	}()
}

func (t *Telemetry) processTelemetryData(wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range t.dataCh {
		go t.sendToDataProcessingService(data)
	}
}

func (t *Telemetry) sendToDataProcessingService(data *vehicle.Telemetry) {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := vehicle.NewDataProcessingServiceClient(conn)
	if _, err := client.ProcessTelemetryData(context.Background(), data); err != nil {
		log.Fatalf("failed to send: %v", err)
	}
}
