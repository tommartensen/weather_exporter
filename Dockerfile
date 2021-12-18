FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .

COPY pkg pkg

RUN go build -o /weather_exporter .

FROM alpine as app

COPY dist dist
COPY --from=builder /weather_exporter .

CMD [ "./weather_exporter" ]
