package main

import "fmt"

func main() {
	var a = 1
	swap(a)
	fmt.Println(a)
}

func swap(x int) {
	x = 10
	fmt.Println(x)
}
