package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(500 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	// game on
	table <- new(Ball) // toss the ball
	time.Sleep(time.Second * 1)
	<-table // game over ,grab the ball

}
