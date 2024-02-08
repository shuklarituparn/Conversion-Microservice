package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/shuklarituparn/Conversion-Microservice/configs"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
)

var (
	Tracer opentracing.Tracer
	Closer io.Closer
	err    error
)

func init() {
	cfg := configs.JaegarConfig()
	Tracer, Closer, err = cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("Failed to create Jaeger tracer: %v", err)
	}
	defer func(closer io.Closer) {
		err := closer.Close()
		if err != nil {

		}
	}(Closer)
	opentracing.SetGlobalTracer(Tracer)
}

func TracingMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		spanCtx, _ := Tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(context.Request.Header))
		span := Tracer.StartSpan(context.Request.URL.Path, ext.RPCServerOption(spanCtx))
		defer span.Finish()

		context.Set("span", span)

		context.Next()

	}
}
