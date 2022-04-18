package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()

	tSaved := time.Now()

	test1(client)

	fmt.Printf("elapse: %v\n", time.Now().Sub(tSaved))

	tSaved = time.Now()

	test2(client)

	fmt.Printf("elapse: %v\n", time.Now().Sub(tSaved))
}

func test1(client *redis.Client) {
	_, _ = client.Incr("tx_pipeline_counter").Result()
	_ = client.Expire("tx_pipeline_counter", time.Hour)
}

func test2(client *redis.Client) {
	pipe := client.TxPipeline()

	_ = pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	pipe.Exec()
}
