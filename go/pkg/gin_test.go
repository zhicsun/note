package pkg

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

var (
	ginAppName = "gin"
	ginPort    = ":8686"
)

func TestGin(t *testing.T) {
	ginStart(ginAppName, ginPort, ginRoute)
}

func ginStart(appName, port string, route func(r *gin.Engine)) {
	r := gin.New()

	r.Use(otelgin.Middleware(appName))
	route(r)

	svc := &http.Server{
		Addr:    port,
		Handler: r,
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

func ginRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		ctx := c.Request.Context()
		fmt.Println(trace.SpanContextFromContext(ctx).TraceID().String())
		fmt.Println(trace.SpanContextFromContext(ctx).SpanID().String())

		restyFormData(ctx)
		ormFind(ctx)
		redisGetSet(ctx)
		getKafkaSyncSend(ctx, brokers, kafkaConfig, topic)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
