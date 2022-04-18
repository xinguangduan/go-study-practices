package main

import (
	"fmt"
)

// select 具有监控多个channel的功能

func main() {
	c := make(chan int)
	quit := make(chan int)
	go calculate(c, quit)
	outputNumber(c, quit)
}

func calculate(c chan int, quit chan int) {
	for i := 0; i < 15; i++ {
		c <- i
	}
	quit <- 0
	fmt.Println("calculate completed.")
}

func outputNumber(c chan int, q chan int) {
	for {
		select {
		case x := <-c:
			fmt.Println(x)
		case <-q:
			fmt.Println("quit")
			return
		}
	}
}
