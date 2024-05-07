package pkg

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestTypeConvert(t *testing.T) {
	fmt.Println(cast.ToString("mayonegg"))
	fmt.Println(cast.ToString(8))
	fmt.Println(cast.ToString(8.31))
	fmt.Println(cast.ToString([]byte("one time")))
	fmt.Println(cast.ToString(nil))
	var foo interface{} = "one more time"
	fmt.Println(cast.ToString(foo))
}
