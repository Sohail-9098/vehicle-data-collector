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
	return loadMQTTConfig()
}

func loadMQTTConfig() (*MQTTConfig, error) {
	bytes, err := os.ReadFile(CONFIG_FILE_PATH)
	if err != nil {
		errStr := fmt.Sprintf("failed to open MQTT config file: %v", err)
		return nil, errors.New(errStr)
	}
	var config *MQTTConfig
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		str := fmt.Sprintf("failed to read MQTT config file: %v", err)
		return nil, errors.New(str)
	}
	return config, nil
}
