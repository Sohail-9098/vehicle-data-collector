package telemetry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTelemetry_NewTelemetry(t *testing.T) {
	atm := NewTelemetry()
	require.NotNil(t, atm)
}
