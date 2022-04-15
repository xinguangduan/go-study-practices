package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("1 -----")
	var numbers []int
	outSlice("numbers", numbers)

	numbers = append(numbers, 0)
	outSlice("numbers", numbers)

	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8)
	outSlice("numbers", numbers)

	n1 := []int{100, 200, 300, 400}
	numbers = append(numbers, n1...)
	outSlice("numbers", numbers)

	numbers = numbers[1:]
	outSlice("numbers", numbers)

	numbers = numbers[:len(numbers)-1]
	outSlice("numbers", numbers)

	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	count := copy(numbers1, numbers)
	outSlice(strconv.Itoa(count), numbers1)

	numbers1[0] = 9999
	numbers[0] = 100000
	outSlice("numbers1", numbers1)
	outSlice("numbers", numbers)
}

func outSlice(name string, x []int) {
	fmt.Printf("content %v \t len=%d \t cap=%d \t value=%v \n", x, len(x), cap(x), name)
}
