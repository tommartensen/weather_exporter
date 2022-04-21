.PHONY: build

VERSION=1.0.0
IMAGE=tommartensen/weather_exporter:$(VERSION)

build:
	go build .

run:
	go run main.go

format: ## Auto-format the code to conform with common Go style
	go fmt ./...

lint: ## Run the linter to enforce best practices
	go vet ./...

docker-build:
	docker build . -t $(IMAGE)

docker-push:
	docker push $(IMAGE)

docker-run:
	docker run -p 9966:9966 $(IMAGE)

helm-deploy-local:
	helm upgrade weatherexporter deploy/chart/ -f deploy/chart/values.yaml -f deploy/chart/local.yaml -f deploy/chart/secret-values.yaml

helm-deploy:
	helm upgrade --install weatherexporter deploy/chart/ -f deploy/chart/values.yaml -f deploy/chart/secret-values.yaml

helm-undeploy:
	helm delete weatherexporter
