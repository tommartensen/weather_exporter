package weatherapi

import (
	"fmt"

	"github.com/tommartensen/weather_exporter/pkg/config"
)

func KelvinToCelsius(temperature float32) float32 {
	return temperature - 273.15
}

func FormatWeather(city config.City, weatherReport weatherReport) string {
	return fmt.Sprintf(
		"weather_temperature{city=\"%s\"} %.1f\n",
		city.Name,
		KelvinToCelsius(weatherReport.Main.Temperature),
	)
}
