package main

import "fmt"

func main() {
	fmt.Println("main goroutine")
	c := make(chan string)
	go func() {
		defer fmt.Println("sub goroutine completed")
		fmt.Println("sub goroutine running...")
		c <- "hello channel"
	}()
	msg := <-c
	close(c)
	fmt.Println("main goroutine completed,msg:", msg)
}
