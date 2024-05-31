package pkg

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	c := cron.New(cron.WithSeconds())

	jobId, err := c.AddJob("*/5 * * * * *", Job1{})
	t.Log(jobId, err)

	jobId, err = c.AddJob("*/10 * * * * *", Job2{})
	t.Log(jobId, err)

	jobId, err = c.AddFunc("*/5 * * * * *", func() {
		fmt.Println(time.Now(), "job3")
	})
	t.Log(jobId, err)

	c.Start()
	<-ch
	c.Stop()
}

type Job1 struct{}

func (t Job1) Run() {
	fmt.Println(time.Now(), "job1")
}

type Job2 struct{}

func (t Job2) Run() {
	fmt.Println(time.Now(), "job2")
}
