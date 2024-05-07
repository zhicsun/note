package pkg

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"testing"
)

func TestJsonQuery(t *testing.T) {
	const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
	name := gojsonq.New().FromString(json).Find("name.first")
	fmt.Println(name.(string))
}
