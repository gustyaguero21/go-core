package metrics

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/zorkian/go-datadog-api.v2"
)

type DatadogMetrics struct {
	client *datadog.Client
}

func NewMetric(apiKey string) *DatadogMetrics {
	client := datadog.NewClient(apiKey, "")
	return &DatadogMetrics{
		client: client,
	}
}

func (d *DatadogMetrics) RecordMetric(metricName string, statusCode int, duration time.Duration) {
	now := time.Now().Unix()
	nowFloat := float64(now)
	durationFloat := duration.Seconds()

	metric := datadog.Metric{
		Metric: &metricName,
		Points: []datadog.DataPoint{
			{&nowFloat, &durationFloat},
		},
		Tags: []string{
			"status_code:" + fmt.Sprintf("%d", statusCode),
		},
		Type: datadog.String("gauge"),
	}

	err := d.client.PostMetrics([]datadog.Metric{metric})
	if err != nil {
		log.Printf("Error sending metrics to Datadog: %v", err)
	}
}
