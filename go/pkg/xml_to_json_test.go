package pkg

import (
	"fmt"
	xj "github.com/basgys/goxml2json"
	"strings"
	"testing"
)

func TestXMLToJson(t *testing.T) {
	xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?><hello>world</hello>`)
	json, err := xj.Convert(xml)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(json.String())
}
