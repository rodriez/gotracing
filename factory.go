package tracing

type TracingProvider interface {
	Setup(traceId string)
	Close()
}

func Build(name string) TracingProvider {
	switch name {
	case "xray":
		return BuildXRayTP()
	case "jaeger":
		return BuildJaegerTP()
	default:
		return BuildConsoleTP()
	}
}
