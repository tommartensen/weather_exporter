package exporter

import (
	"fmt"
	"net/http"

	"github.com/tommartensen/weather_exporter/pkg/config"
	"github.com/tommartensen/weather_exporter/pkg/weatherapi"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	config := config.LoadConfig()

	openWeatherClient := weatherapi.GetClient()
	for _, city := range config.Cities {
		weatherReport := weatherapi.GetCurrentWeather(openWeatherClient, city)
		writeWeatherReport(w, city, weatherReport)
	}
}

func gatherRelevantMetrics(weatherReport weatherapi.WeatherReport) map[string]interface{} {
	return map[string]interface{}{
		"temperature":           KelvinToCelsius(weatherReport.Main.Temperature),
		"temperature_perceived": KelvinToCelsius(weatherReport.Main.TemperaturePerceived),
		"humidity":              weatherReport.Main.Humidity,
		"pressure":              weatherReport.Main.Pressure,
		"wind_speed":            weatherReport.Wind.Speed,
		"wind_degree":           weatherReport.Wind.Degree,
		"cloud_cover":           weatherReport.CloudCover.All,
	}
}

func writeWeatherReport(w http.ResponseWriter, city config.City, weatherReport weatherapi.WeatherReport) {
	relevantMetrics := gatherRelevantMetrics(weatherReport)

	output := ""
	for metric, value := range relevantMetrics {
		output += FormatMetric(metric, city, value)
	}
	fmt.Fprint(w, output)
}
