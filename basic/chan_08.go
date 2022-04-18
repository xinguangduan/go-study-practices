package main

import "fmt"
import "time"

// 通道同步
func worker(done chan bool, ts time.Duration) {
	fmt.Print("working...")
	time.Sleep(ts)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool, 1)
	go worker(done, time.Second*3)
	<-done
	fmt.Println("work1 completed.")
	subTask := make(chan bool, 1)
	go worker(subTask, time.Second*5)
	<-subTask
	fmt.Println("sub task completed.")
}
