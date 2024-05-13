package pkg

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestRedis(t *testing.T) {
	redisGetSet(context.Background())
}

func redisGetSet(ctx context.Context) {
	client, err := getRedisConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client.Set(ctx, "key", "value", 0).Result())
	fmt.Println(client.Get(ctx, "key").Result())
	fmt.Println(client.Del(ctx, "key").Result())
}

func getRedisConn() (*redis.ClusterClient, error) {
	addrs := []string{
		"web-redis-cache-01.db.sit13.dom:52001",
		"web-redis-cache-02.db.sit13.dom:52001",
		"web-redis-cache-03.db.sit13.dom:52001",
	}
	opt := &redis.ClusterOptions{
		Addrs:        addrs,
		Password:     "ee@redis^#**",
		PoolSize:     100,
		MinIdleConns: 50,
	}

	client := redis.NewClusterClient(opt)

	if err := redisotel.InstrumentTracing(client); err != nil {
		return nil, err
	}

	return client, nil
}
