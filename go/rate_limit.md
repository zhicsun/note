# 速率限制

## 固定速率

```go
package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limiter := rate.NewLimiter(rate.Limit(100), 1) // 允许每秒100次
	for i := 0; i < 200; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}

		go process()
	}
}

func process() {
	fmt.Println("Request processed successfully.")
	time.Sleep(time.Millisecond)
}

```

## 令牌桶

```go
package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limiter := rate.NewLimiter(rate.Limit(10), 5)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond)
	defer cancel()
	for i := 0; i < 200; i++ {
		if err := limiter.Wait(ctx); err != nil {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}

		go process()
	}
}

func process() {
	fmt.Println("Request processed successfully.")
	time.Sleep(time.Millisecond)
}

```

## 动态速率

```go
package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limiter := rate.NewLimiter(rate.Limit(10), 1)

	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("---adjust limiter---")
		limiter.SetLimit(rate.Limit(200)) // 将 limiter 提升到每秒 200
	}()

	for i := 0; i < 3000; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			time.Sleep(time.Millisecond * 100)
			continue
		}

		process()
	}
}

func process() {
	fmt.Println("Request processed successfully.")
	time.Sleep(time.Millisecond * 10)
}

```

## 自适应速率

```go
package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	limiter := rate.NewLimiter(rate.Limit(10), 1)

	go func() {
		for {
			time.Sleep(time.Second * 10)
			responseTime := measureResponseTime()
			if responseTime > 500*time.Millisecond {
				fmt.Println("---adjust limiter 50---")
				limiter.SetLimit(rate.Limit(50))
			} else {
				fmt.Println("---adjust limiter 100---")
				limiter.SetLimit(rate.Limit(100))
			}
		}
	}()

	for i := 0; i < 3000; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			time.Sleep(time.Millisecond * 100)
			continue
		}
		process()
	}
}

func measureResponseTime() time.Duration {
	return time.Millisecond * 100
}

func process() {
	fmt.Println("Request processed successfully.")
	time.Sleep(time.Millisecond * 10)
}

```