module github.com/rodriez/gotracing

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	go.opentelemetry.io/contrib/propagators/aws v0.22.0
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0-RC2
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/grpc v1.39.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
