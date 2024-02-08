package configs

import "github.com/uber/jaeger-client-go/config"

var Jaegarconf *config.Configuration

func init() {
	Jaegarconf = &config.Configuration{
		ServiceName: "conversion-service",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
}

func JaegarConfig() *config.Configuration {
	return Jaegarconf
}
