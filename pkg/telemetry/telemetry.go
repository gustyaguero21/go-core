package telemetry

import (
	"log"

	"github.com/DataDog/datadog-go/statsd"
)

var client *statsd.Client

func init() {
	var err error
	client, err = statsd.New("localhost:8125")
	if err != nil {
		log.Fatal("Error initializing Datadog client:", err)
	}
}

func Incr(metricName string, statusCode int) {
	metricNameWithStatus := metricName + ".statuscode." + string(rune(statusCode))

	if err := client.Incr(metricNameWithStatus, nil, 1); err != nil {
		log.Println("Error incrementing metric:", err)
	}
}
