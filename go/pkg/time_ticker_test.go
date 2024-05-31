package pkg

import (
	"context"
	"testing"
	"time"
)

func TestTimeTicker(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case now := <-ticker.C:
			t.Log(now)
		case <-ctx.Done():
			goto done
		}
	}

done:
	t.Log("done")
}
