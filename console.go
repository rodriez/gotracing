package tracing

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type ConsoleTP struct {
}

func BuildConsoleTP() *ConsoleTP {
	return &ConsoleTP{}
}

func (ctp *ConsoleTP) Setup(traceId string) {
	exporter, err := stdouttrace.New(stdouttrace.WithWriter(log.Writer()))
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSyncer(exporter),
	)

	otel.SetTracerProvider(tp)
	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	otel.SetTextMapPropagator(propagator)
}

func (ctp *ConsoleTP) Close() {
	log.Println("Trace provider closed")
}
