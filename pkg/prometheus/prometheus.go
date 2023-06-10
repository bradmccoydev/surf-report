package prometheus

import (
	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func pushGaugeMetric(pushGatewayURL string, metricName string, metricValue float64) {
	// Create a new Gauge metric for swell direction
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: metricName,
		Help: "Custom metric",
	})

	// Set the value for swell direction
	gauge.Set(metricValue)

	// Create a new Pusher and specify the Pushgateway URL
	pusher := push.New(pushGatewayURL, "surf")

	// Push the swell direction metric to the Pushgateway
	err := pusher.Collector(gauge)
	if err != nil {
		log.Fatal("Failed to push metric to Pushgateway:", err)
	}

	// Push the metrics to the Pushgateway
	pushError := pusher.Push()
	if pushError != nil {
		log.Fatal("Failed to push metrics to Pushgateway:", err)
	}

	fmt.Println("Metrics pushed successfully")
}
