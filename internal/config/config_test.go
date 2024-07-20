package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_NewMQTTConfig(t *testing.T) {
	_, err := NewMQTTConfig()
	require.NoError(t, err)
}
