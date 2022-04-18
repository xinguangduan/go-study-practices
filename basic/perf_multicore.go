package main

import (
	"fmt"
	"time"
)

var count = 0

func goFunc(i int) {
	fmt.Println("go routine ", i, "...")
	time.Sleep(time.Microsecond)
	count++
}

func main() {
	var beginTime = time.Now()
	for i := 0; i < 100; i++ {
		go goFunc(i) // open an concurrent coroutines
	}
	time.Sleep(time.Minute)
	fmt.Println(count)
	fmt.Println(time.Now().Sub(beginTime))
}
