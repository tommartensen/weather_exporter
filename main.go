package main

import (
	"log"
	"net/http"

	"github.com/tommartensen/weather_exporter/pkg/exporter"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.HandleFunc("/healthz", exporter.Healthcheck)
	http.HandleFunc("/metrics", exporter.Serve)

	log.Fatal(http.ListenAndServe(":9966", nil))
}
