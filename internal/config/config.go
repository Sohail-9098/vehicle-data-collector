package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	CONFIG_FILE_PATH = "../../configs/mqtt_config.yaml"
)

type MQTTConfig struct {
	MQTT struct {
		Broker   string `yaml:"broker"`
		ClientID string `yaml:"client_id"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mqtt"`
}

func NewMQTTConfig() (*MQTTConfig, error) {
	return loadMQTTConfig(CONFIG_FILE_PATH)
}

func loadMQTTConfig(filepath string) (*MQTTConfig, error) {
	var mqttConfig *MQTTConfig
	bytes, err := readMQTTConfigFile(filepath)
	if err != nil {
		return nil, err
	}
	if err := decodeMQTTConfig(bytes, &mqttConfig); err != nil {
		return nil, err
	}
	return mqttConfig, nil
}

func readMQTTConfigFile(filepath string) ([]byte, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		errStr := fmt.Sprintf("failed to read MQTT config file: %v", err)
		return nil, errors.New(errStr)
	}
	return bytes, nil
}

func decodeMQTTConfig(bytes []byte, mqttConfig **MQTTConfig) error {
	if err := yaml.Unmarshal(bytes, &mqttConfig); err != nil {
		str := fmt.Sprintf("failed to decode MQTT config file: %v", err)
		return errors.New(str)
	}
	return nil
}
