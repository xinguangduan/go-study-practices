package main

import "fmt"

func main() {
	c := make(chan int, 3)

	go func() {
		for i := 0; i < cap(c); i++ {
			c <- i
		}
		close(c)
	}()
	for d := range c {
		fmt.Println(d)
	}
	fmt.Println("main goroutine completed")
}
