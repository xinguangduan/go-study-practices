package main

import (
	"fmt"
)

// Go 的通道选择器 让你可以同时等待多个通道操作。Go 协程和通道以及选择器的结合是 Go 的一个强大特性。
func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	//各个通道将在若干时间后接收一个值，这个用来模拟例如并行的 Go 协程中阻塞的 RPC 操作
	go func() {
		//time.Sleep(time.Second * 1)
		for i := 0; i < 10; i++ {
			fmt.Println("task one ", i)
		}
		c1 <- "one"
	}()
	go func() {
		//time.Sleep(time.Second * 2)
		for i := 0; i < 10; i++ {
			fmt.Println("task two ", i)
		}
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
