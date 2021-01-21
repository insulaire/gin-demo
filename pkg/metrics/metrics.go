package metrics

import (
	"fmt"
	"gin-demo/global"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	labels = []string{"status", "endpoint", "method"}

	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: global.ServiceSetting.Name,
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, labels,
	)
)

func init() {
	prometheus.Register(reqCount)
}

func Inc(c *gin.Context) {
	c.Next()
	status := fmt.Sprintf("%d", c.Writer.Status())
	endpoint := c.Request.URL.Path
	method := c.Request.Method

	lvs := []string{status, endpoint, method}
	reqCount.WithLabelValues(lvs...).Inc()
}
