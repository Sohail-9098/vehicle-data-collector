package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"gopkg.in/yaml.v3"
)

const (
	CONFIG_FILE_PATH = "../../configs/mqtt_config.yaml"
)

type MQTTConfig struct {
	MQTT struct {
		Broker   string `json:"broker"`
		ClientID string `json:"client_id"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"mqtt"`
}

func New() (*MQTTConfig, error) {
	config := &MQTTConfig{}
	configStr, err := getSecretValue("vehicle/vehicle-data-collector/mqtt/credentials")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(configStr), &config.MQTT); err != nil {
		return nil, errors.New("failed to decode config JSON")
	}
	config.MQTT.Broker = "tcp://localhost:1883"
	return config, nil
}

func getSecretValue(secretName string) (string, error) {
	// load aws config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		errStr := fmt.Sprintf("failed to load config: %v", err)
		return "", errors.New(errStr)
	}

	// load 'secretsmanager' service with help of config
	svc := secretsmanager.NewFromConfig(cfg)
	resp, err := svc.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	})
	if err != nil {
		errStr := fmt.Sprintf("failed to get secret value: %v", err)
		return "", errors.New(errStr)
	}
	return *resp.SecretString, nil
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
