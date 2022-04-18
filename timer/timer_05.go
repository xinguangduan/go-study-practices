package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	done := make(chan bool)
	go func() {
		<-timer1.C
		fmt.Println("Timer 1 fired")
		done <- true
	}()
	<-done
}
