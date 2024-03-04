package prometheus

import "github.com/prometheus/client_golang/prometheus"

var (
	ConvertApiPingCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Convert_API_Ping_Counter",
		Help: "Number of pings made to the Convert API endpoint",
	})

	CutApiPingCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Cut_API_Ping_Counter",
		Help: "Number of pings made to the CUT API endpoint",
	})
	WatermarkApiPingCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Watermark_API_Ping_Counter",
		Help: "Number of pings made to the Watermark API endpoint",
	})
	ScreenshotApiPingCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Screenshot_API_Ping_Counter",
		Help: "Number of pings made to the Screenshot API endpoint",
	})
)

func Init() {
	prometheus.MustRegister(ConvertApiPingCounter)
	prometheus.MustRegister(CutApiPingCounter)
	prometheus.MustRegister(WatermarkApiPingCounter)
	prometheus.MustRegister(ScreenshotApiPingCounter)
}
