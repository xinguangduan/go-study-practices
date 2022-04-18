package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 10)
	go func() {
		<-timer.C
		fmt.Println("Timer fired")
	}()
	stopped := timer.Stop()
	if stopped {
		fmt.Println("Timer stopped")
	}
	time.Sleep(2 * time.Second)
}
