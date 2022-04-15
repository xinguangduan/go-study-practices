package main

import (
	"fmt"
	"time"
)

func main() {
	begin := time.Now().UnixNano()
	//var numbers []int //bad performance
	numbers := make([]int, 0, 100) //best performance
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i)
		outSliceData("numbers", numbers)
	}
	fmt.Printf("cost:%d ns\n", time.Now().UnixNano()-begin)

}

func outSliceData(name string, x []int) {
	fmt.Printf("value=%v \t address:%p \t content %v \t len=%d \t cap=%d \n", name, x, x, len(x), cap(x))
}
