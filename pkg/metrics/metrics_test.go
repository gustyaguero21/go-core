package metrics

import (
	"testing"
	"time"
)

func TestRecordMetric(t *testing.T) {
	ddMetrics := NewMetric("fake-api-key")

	metricName := "test_metric"
	statusCode := 200
	duration := time.Duration(2 * time.Second)

	ddMetrics.RecordMetric(metricName, statusCode, duration)

	t.Log("Metric recorded successfully")
}
