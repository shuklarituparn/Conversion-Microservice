package middlewares

import (
	"fmt"
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
	opentracing.SetGlobalTracer(Tracer)
}

func TracingMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		spanCtx, err := Tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(context.Request.Header))
		if err != nil {
			// Handle extraction error
			log.Println("Error extracting span context:", err)
		}

		spanName := fmt.Sprintf("%s %s", context.Request.Method, context.Request.URL.Path)
		span := Tracer.StartSpan(spanName, ext.RPCServerOption(spanCtx))

		defer func() {
			if r := recover(); r != nil {
				// Recover from panic and finish the span
				span.Finish()
				panic(r)
			}
			span.Finish()
		}()

		context.Set("span", span)
		context.Next()
	}
}
