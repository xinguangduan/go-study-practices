package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("current num is ", i)
	}
}
