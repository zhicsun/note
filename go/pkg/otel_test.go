package pkg

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"testing"
)

func TestOtel(t *testing.T) {
	err := startAgent()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stopAgent()
	go ginStart(ginAppName, ginPort, ginRoute)
	kafkaConsumer(context.Background())
}

var (
	tp *sdktrace.TracerProvider
)

func startAgent() error {
	rpcUrl := "10.252.239.234:4317"
	serviceName := "lms"
	env := "dev"
	containerName := "note"
	sampler := 1.0

	grpcOps := []otlptracegrpc.Option{
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(rpcUrl),
	}
	exporter, err := otlptracegrpc.New(context.Background(), grpcOps...)
	if err != nil {
		return err
	}

	resources := []attribute.KeyValue{
		semconv.ServiceName(serviceName),
		semconv.DeploymentEnvironment(env),
		semconv.ContainerName(containerName),
	}
	opts := []sdktrace.TracerProviderOption{
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(sampler))),
		sdktrace.WithResource(resource.NewSchemaless(resources...)),
		sdktrace.WithBatcher(exporter),
	}
	tp = sdktrace.NewTracerProvider(opts...)
	otel.SetTracerProvider(tp)
	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		fmt.Println(err.Error())
	}))
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return nil
}

func stopAgent() {
	err := tp.Shutdown(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
}
