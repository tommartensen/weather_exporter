package weatherapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/tommartensen/weather_exporter/pkg/config"
)

type OpenWeatherClient struct {
	ApiToken string
}

type WeatherReport struct {
	Coordinates struct {
		Longitude float32 `json:"lon"`
		Latitude  float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Type        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temperature          float32 `json:"temp"`
		TemperaturePerceived float32 `json:"feels_like"`
		TemperatureMin       float32 `json:"temp_min"`
		TemperatureMax       float32 `json:"temp_max"`
		Humidity             uint8   `json:"humidity"`
		Pressure             uint16  `json:"pressure"`
		PressureAtSeaLevel   uint16  `json:"sea_level"`
		PressureAtGround     uint16  `json:"grnd_level"`
	} `json:"main"`
	Visibility uint16 `json:"visibility"`
	Wind       struct {
		Speed     float32 `json:"speed"`
		Degree    uint16  `json:"deg"`
		GustSpeed float32 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		NextHour float32 `json:"1h"`
	} `json:"rain"`
	CloudCover struct {
		All uint8 `json:"all"`
	} `json:"clouds"`
	Sun struct {
		SunRise uint32 `json:"sunrise"`
		SunSet  uint32 `json:"sunset"`
	} `json:"sys"`
	Timezone int16 `json:"timezone"`
}

const API_BASE_URL string = "api.openweathermap.org"
const API_SCHEMA string = "https"
const API_VERSION string = "2.5"

func GetClient() *OpenWeatherClient {
	apiToken := config.LoadEnvironment().ApiToken
	openWeatherClient := OpenWeatherClient{ApiToken: apiToken}
	return &openWeatherClient
}

func buildApiUrl(openWeatherClient *OpenWeatherClient, city config.City) string {
	encodedCityName := url.QueryEscape(city.Name)
	return fmt.Sprintf(
		"%s://%s/data/%s/weather?q=%s,%s&appid=%s",
		API_SCHEMA, API_BASE_URL, API_VERSION,
		encodedCityName, city.CountryCode, openWeatherClient.ApiToken,
	)
}

func parseResponse(body []byte) WeatherReport {
	openWeatherApiResponse := WeatherReport{}
	if err := json.Unmarshal(body, &openWeatherApiResponse); err != nil {
		log.Fatal(err)
	}
	return openWeatherApiResponse
}

func GetCurrentWeather(openWeatherClient *OpenWeatherClient, city config.City) WeatherReport {
	url := buildApiUrl(openWeatherClient, city)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Fatalf("Request failed: %s -> %d", url, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	weatherReport := parseResponse(body)
	return weatherReport
}
