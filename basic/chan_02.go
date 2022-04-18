package main

import "fmt"

func task1(c chan string) {
	fmt.Println("task1 is running")
	c <- "sss"
}

func task2(c chan string) {
	fmt.Println("task2 is running")
	msg := <-c
	defer fmt.Println("task2 receive ", msg)
}

func task3(c chan string) {
	fmt.Println("task3 is running")
	defer fmt.Println("task3 completed add message")
	c <- "你好"
}

func main() {
	c := make(chan string)
	go task1(c)
	go task2(c)
	go task3(c)
	m := <-c
	fmt.Println("completed.", m)

}
