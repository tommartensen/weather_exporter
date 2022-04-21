package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type City struct {
	Name        string `yaml:"name"`
	CountryCode string `yaml:"countryCode"`
}

type weatherConfiguration struct {
	Cities      []City `yaml:"cities"`
	DefaultCity City   `yaml:"defaultCity"`
}

type environment struct {
	ApiToken string
}

func LoadEnvironment() environment {
	environ := environment{ApiToken: os.Getenv("OPENWEATHER_API_TOKEN")}
	return environ
}

func LoadConfig() weatherConfiguration {
	filename, _ := filepath.Abs("./config/cities.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config := weatherConfiguration{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}
