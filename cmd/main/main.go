package main

import "github.com/Sohail-9098/vehicle-data-collector/internal/telemetry"

func main() {
	data := telemetry.NewTelemetryData()
	data.FetchAndProcessTelemetryData()
}
