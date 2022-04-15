package main

import "fmt"

func main() {
	a := 10
	var b *int = &a

	fmt.Println(a)
	fmt.Println(b)
	a++
	*b++
	fmt.Println(a)
}
