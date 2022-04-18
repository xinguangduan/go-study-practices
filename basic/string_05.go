package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "able"
	b := "Able"

	s := strings.EqualFold(a, b)
	r := a == b
	fmt.Println(s)
	fmt.Println(r)
}
