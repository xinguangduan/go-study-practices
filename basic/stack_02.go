package main

import "fmt"

func showPoint() *int {
	num := 1
	point := &num
	return point
}
func main() {
	var point *int
	point = showPoint()
	fmt.Println(*point)
}
