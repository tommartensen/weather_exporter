package exporter

import (
	"fmt"
	"net/http"

	"github.com/tommartensen/weather_exporter/pkg/config"
	"github.com/tommartensen/weather_exporter/pkg/weatherapi"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	config.LoadConfig()
	fmt.Fprintf(w, "Ok")
}

func ApiHealthcheck(w http.ResponseWriter, r *http.Request) {
	configuration := config.LoadConfig()
	openWeatherClient := weatherapi.GetClient()

	weatherapi.GetCurrentWeather(openWeatherClient, configuration.DefaultCity)
	fmt.Fprintf(w, "Ok")
}
