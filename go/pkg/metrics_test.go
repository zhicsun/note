package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestMetrics(t *testing.T) {
	startMetrics()
}

var (
	counter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "dev",
			Subsystem: "sms",
			Name:      "http_requests_total",
			Help:      "The total number of handled HTTP requests.",
		},
		[]string{"code", "method"},
	)

	gauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dev",
			Subsystem: "sms",
			Name:      "home_temperature_celsius",
			Help:      "The current temperature in degrees Celsius.",
		},
		[]string{"house", "room"},
	)

	histogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "dev",
			Subsystem: "sms",
			Name:      "http_request_duration_seconds_histogram",
			Help:      "A histogram of the HTTP request durations in seconds.",
			Buckets:   []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"house", "room"},
	)

	summary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: "dev",
			Subsystem: "sms",
			Name:      "http_request_duration_seconds_summary",
			Help:      "A summary of the HTTP request durations in seconds.",
			Objectives: map[float64]float64{
				0.5:  0.05,
				0.9:  0.01,
				0.99: 0.001,
			},
		},
		[]string{"house", "room"},
	)
)

func startMetrics() {
	registry := prometheus.NewRegistry()
	registry.MustRegister(counter, gauge, histogram, summary)

	svcMux := http.NewServeMux()
	svcMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		counter.With(prometheus.Labels{"code": "200", "method": "get"}).Inc()
		counter.With(prometheus.Labels{"code": "200", "method": "post"}).Inc()

		gauge.With(prometheus.Labels{"house": "1", "room": "living"}).Set(19.5)
		gauge.With(prometheus.Labels{"house": "1", "room": "bedroom"}).Set(22.3)
		gauge.With(prometheus.Labels{"house": "2", "room": "living"}).Set(20.0)
		gauge.With(prometheus.Labels{"house": "2", "room": "bedroom"}).Set(21.0)

		histogram.WithLabelValues("1", "living").Observe(2)
		histogram.WithLabelValues("1", "living").Observe(0.5)
		histogram.WithLabelValues("1", "living").Observe(0.5)

		summary.WithLabelValues("1", "living").Observe(0.1)
		summary.WithLabelValues("1", "living").Observe(0.2)
		summary.WithLabelValues("1", "living").Observe(0.3)
		summary.WithLabelValues("1", "living").Observe(0.4)
	})
	svcMux.HandleFunc("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{EnableOpenMetrics: true}).ServeHTTP)
	svcMux.HandleFunc("/debug/pprof/", pprof.Index)
	svcMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	svcMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	svcMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	svcMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	port := ":8688"
	svc := http.Server{
		Addr:    port,
		Handler: svcMux,
	}
	go func() {
		if err := svc.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	timeout := time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := svc.Shutdown(ctx); err != nil {
		fmt.Println(err.Error())
	}
}
