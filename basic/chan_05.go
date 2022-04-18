package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)
	fmt.Println("chan len=", len(c), ",cap=", cap(c))
	go taskA(c)
	time.Sleep(2 * time.Second)
	go taskB(c)
	time.Sleep(20 * time.Second)

}

func taskA(c chan int) {

	defer fmt.Println("taskA add data to chan completed,", len(c))
	for i := 0; i < cap(c); i++ {
		c <- i
		fmt.Println("taskA ", i, "len=", len(c))
	}
}
func taskB(c chan int) {
	defer fmt.Println("taskB fetch data completed.", len(c))
	for i := 0; i < cap(c); i++ {
		n := <-c
		fmt.Println("taskB ", n, "len:", len(c))
	}
}
