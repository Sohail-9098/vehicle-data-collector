package mqtt

import (
	"testing"

	"github.com/Sohail-9098/vehicle-data-collector/internal/config"
	"github.com/stretchr/testify/require"
)

func TestMQTT_NewClient(t *testing.T) {
	config, err := config.NewMQTTConfig()
	require.NoError(t, err, "error reading config file")
	client := NewClient(config.MQTT.Broker, config.MQTT.ClientID, config.MQTT.Username, config.MQTT.Password)
	require.NotNil(t, client)
}

func TestMQTT_Connect_Invalid(t *testing.T) {
	testConfig := config.MQTTConfig{}
	testConfig.MQTT.Broker = "testBroker"
	testConfig.MQTT.ClientID = "testClientID"
	testConfig.MQTT.Username = "fakeUsername"
	testConfig.MQTT.Password = "fakePassword"
	client := NewClient(testConfig.MQTT.Broker, testConfig.MQTT.ClientID, testConfig.MQTT.Username, testConfig.MQTT.Password)
	err := client.Connect()
	require.Error(t, err)
}

func TestMQTT_Connect_Valid(t *testing.T) {
	config, _ := config.NewMQTTConfig()
	client := NewClient(config.MQTT.Broker, config.MQTT.ClientID, config.MQTT.Username, config.MQTT.Password)
	err := client.Connect()
	require.NoError(t, err)
	client.Disconnect()
}
