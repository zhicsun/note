package paradigm

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	c := &Consumer{}
	c.quit = make(chan os.Signal, 1)
	signal.Notify(c.quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for {
			err := c.Send(Msg{Content: time.Now().String()})
			if err != nil {
				t.Log(err)
				return
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("消费者 %d", i)
		go func() {
			defer wg.Done()
			msg := c.Subscribe(100)
			for v := range msg {
				fmt.Println(name, v.Content)
			}
		}()
	}
	wg.Wait()
}

type Msg struct {
	Content string
}

type Consumer struct {
	chs  []chan Msg
	quit chan os.Signal
}

func (c *Consumer) Send(m Msg) error {
	for _, ch := range c.chs {
		select {
		case ch <- m:
		case <-c.quit:
			c.Close()
			return nil
		default:
			return errors.New("消息队列已满")
		}
	}

	return nil
}

func (c *Consumer) Subscribe(capacity int) <-chan Msg {
	res := make(chan Msg, capacity)
	c.chs = append(c.chs, res)
	return res
}

func (c *Consumer) Close() {
	chs := c.chs
	c.chs = nil

	for _, ch := range chs {
		close(ch)
	}
}
