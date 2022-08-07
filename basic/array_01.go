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
	p := fmt.Println
	var xx []int
	xx = make([]int, 20)
	y := make([]int, 30)
	p(y)
	p(xx)
	p(xx)

}
