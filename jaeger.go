package gotracing

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type JaegerTP struct {
	exporter *jaeger.Exporter
	tp       *sdktrace.TracerProvider
}

func BuildJaegerTP() *JaegerTP {
	return &JaegerTP{}
}

func (jtp *JaegerTP) Setup(traceId string) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		log.Fatal(err)
	}

	jtp.exporter = exp

	jtp.tp = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSyncer(jtp.exporter),
	)

	otel.SetTracerProvider(jtp.tp)
}

func (tp *JaegerTP) Close() {
	log.Println("Trace provider closed")
}
