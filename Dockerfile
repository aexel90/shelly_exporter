FROM golang:alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /shelly_exporter

FROM alpine:latest
WORKDIR /
COPY --from=build /shelly_exporter /shelly_exporter
COPY shelly-metrics.json ./
EXPOSE 9773

ENTRYPOINT [ "sh", "-c", "/shelly_exporter -metrics-file shelly-metrics.json" ]