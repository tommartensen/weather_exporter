package exporter

import (
	"fmt"

	"github.com/tommartensen/weather_exporter/pkg/config"
)

func KelvinToCelsius(temperature float32) float32 {
	return temperature - 273.15
}

func FormatMetric(name string, city config.City, value interface{}) string {
	return fmt.Sprintf(
		"weather_%s{city=\"%s\"} %.2v\n",
		name, city.Name, value,
	)
}
