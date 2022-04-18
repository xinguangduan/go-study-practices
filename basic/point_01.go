package main

import "fmt"

func main() {
	a := 100
	b := 222
	var p *int
	switchValue(&a, &b)
	fmt.Println(a, b)
	fmt.Println(&p)
}

func switchValue(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}
