package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gustyaguero21/go-core/pkg/metrics"
)

type Error struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}

func (er *Error) Error() string {
	return fmt.Sprintf("Status: %d, Error: %s", er.Status, er.Err)
}

func NewError(ctx *gin.Context, status int, err string, ddMetrics *metrics.DatadogMetrics, metricName string) {
	errorStruct := &Error{
		Status: status,
		Err:    err,
	}

	ddMetrics.RecordMetric(metricName, status, 0)

	ctx.JSON(status, errorStruct)
}
