package main

import "fmt"

func main() {
	a := 100
	b := 200
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println(a, b)

	swap1(&a, &b)
	fmt.Println(a, b)

}
func swap(a, b int) (int, int) {
	return b, a
}

func swap2(a, b *int) (*int, *int) {
	return b, a
}

func swap1(a, b *int) {
	*a, *b = *b, *a
}
