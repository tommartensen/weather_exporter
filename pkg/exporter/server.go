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
		fmt.Fprint(w, weatherapi.FormatWeather(city, weatherReport))
	}
}
