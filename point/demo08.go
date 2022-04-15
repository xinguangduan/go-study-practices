package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	i := 0
	for i = 0; i < 10; i++ {
		//fmt.Println(i)
	}
	cost := time.Since(start)
	fmt.Printf("%v\n", cost)
	fmt.Printf("%v", time.Now().Format("2006-01-02 15:04:05"))

}
