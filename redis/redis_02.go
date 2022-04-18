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

	pipe := client.TxPipeline()

	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	// Execute
	//
	//     MULTI
	//     INCR pipeline_counter
	//     EXPIRE pipeline_counts 3600
	//     EXEC
	//
	// using one redisdb-server roundtrip.
	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}
