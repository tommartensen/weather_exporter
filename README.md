# Weather Exporter

This exporter queries Openweather API for a given list of cities and exposes the weather information in Prometheus metric format for scraping.
It comes with a Helm chart to deploy into K8s

## Building

1. `make docker-build`
2. `make docker-push`

## Deployment

1. Get an Openweather API token
2. Set the token under the `openweatherApiToken` value in `deploy/chart/values.yaml` or provide another values file to the install command.
3. `helm upgrade --install weatherexporter deploy/chart/ -f deploy/chart/values.yaml`

The `local.yaml` and `values.yaml` come with a non-existent token to let the container start. 

The service type can be set through the `service.type` value, e.g. to `NodePort` or `ClusterIP`.
The `local.yaml` is an example for this.
The Ingress is disabled by default, which be controlled in the `ingress.enabled` and `ingress.host` settings.  
