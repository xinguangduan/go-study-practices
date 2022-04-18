package main

import (
	"fmt"
)

func main() {
	ft := fmt.Println
	var x [58]int
	ft(x)
	ft(len(x))
	ft(x[42])
	x[42] = 777
	ft(x[42])

}
