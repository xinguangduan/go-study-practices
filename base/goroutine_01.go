package main

import (
	"fmt"
	"time"
)

func newTask() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go newTask()
	fmt.Println("main thread")
	time.Sleep(10 * time.Second)
}
