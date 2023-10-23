# 功能选项模式
```go
package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	_ "net/http/pprof"
	"time"
)

func main() {
	srv := NewServer("127.0.0.1", 3306, Protocol("udp"), Timeout(60), MaxConn(80), TLS(nil))
	data, _ := json.Marshal(srv)
	fmt.Println(string(data))
}

func NewServer(addr string, port int, options ...func(*Server)) *Server {
	srv := Server{
		Addr: addr,
		Port: port,
		Config: &Config{
			Protocol: "tcp",
			Timeout:  30,
			MaxConn:  10,
			TLS:      nil,
		},
	}

	for _, option := range options {
		option(&srv)
	}
	
	return &srv
}

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConn(maxConn int) Option {
	return func(s *Server) {
		s.MaxConn = maxConn
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

type Option func(*Server)

type Server struct {
	Addr string
	Port int
	*Config
}

type Config struct {
	Protocol string
	Timeout  time.Duration
	MaxConn  int
	TLS      *tls.Config
}

```