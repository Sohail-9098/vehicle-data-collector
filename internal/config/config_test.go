package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_NewMQTTConfig_Valid(t *testing.T) {
	config, err := NewMQTTConfig()
	require.Equal(t, config.MQTT.ClientID, "vehicle_data_collector")
	require.NoError(t, err)
}

func TestConfig_LoadConfig_Invalid(t *testing.T) {
	_, err := loadMQTTConfig("invalidpath")
	require.Error(t, err)
}

func TestConfig_ReadMQTTConfig_Invalid(t *testing.T) {
	_, err := readMQTTConfigFile("abc")
	require.Error(t, err)
}

func TestConfig_ReadMQTTConfig_Valid(t *testing.T) {
	_, err := readMQTTConfigFile(CONFIG_FILE_PATH)
	require.NoError(t, err)
}

func TestConfig_DecodeMQTTConfig_Valid(t *testing.T) {
	bytes, err := readMQTTConfigFile(CONFIG_FILE_PATH)
	require.NoError(t, err)
	var mqttConfig *MQTTConfig
	err = decodeMQTTConfig(bytes, &mqttConfig)
	require.NoError(t, err)
}

func TestConfig_DecodeMQTTConfig_Invalid(t *testing.T) {
	var mqttConfig *MQTTConfig
	err := decodeMQTTConfig([]byte{'x'}, &mqttConfig)
	require.Error(t, err)
}
