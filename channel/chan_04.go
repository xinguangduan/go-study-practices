package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	go func() {
		time.Sleep(time.Second * 1)
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Printf("add value to channel %d\n", len(c))
		}
	}()

	v := <-c
	fmt.Println(v)

	fmt.Println("finished.")
}
