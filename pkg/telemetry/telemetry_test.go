package telemetry

import (
	"log"
	"testing"

	"github.com/DataDog/datadog-go/statsd"
)

func initTestClient() {
	var err error
	client, err = statsd.New("localhost:8125")
	if err != nil {
		log.Fatal("Error initializing Datadog client:", err)
	}
}

func TestIncr(t *testing.T) {
	initTestClient()

	Incr("http.internal_error", 500)

	t.Log("Métrica enviada correctamente a Datadog.")
}
