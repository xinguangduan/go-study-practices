package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	go func() {
		defer close(c)
		defer fmt.Println("goroutine over.")
		c <- 2222
	}()
	//time.Sleep(3 * time.Second)
	//defer close(c)
	for {
		if num, ok := <-c; ok {
			fmt.Println("number=", num)
		}
		break
	}
	fmt.Println("main goroutine over")

}
