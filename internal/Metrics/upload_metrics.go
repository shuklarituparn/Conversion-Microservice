package Metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	UploadKafkaMetrics = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Upload_kafka_metrics",
		Help: "Metrics to track the upload messages produced by kafka",
	})
	ConvertKafkaMetrics = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Upload_kafka_metrics",
		Help: "Metrics to track the convert messages produced by kafka",
	})
)

func init() {
	prometheus.MustRegister(UploadKafkaMetrics)
	prometheus.MustRegister(ConvertKafkaMetrics)
}
