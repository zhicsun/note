package pkg

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
	"testing"
)

func TestResty(t *testing.T) {
	RestyFormData(context.Background())
}

func RestyFormData(ctx context.Context) {
	path := "http://sms.k8s.sit13.dom/v1/send/sms/verification/code"
	var res any
	client := resty.NewWithClient(&http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	})
	response, err := client.R().SetContext(ctx).SetFormData(map[string]string{
		"appId": "f5dc10e438e35c6d8dc1a3e5c2b521b4",
	}).SetResult(&res).Post(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response, res)
}
