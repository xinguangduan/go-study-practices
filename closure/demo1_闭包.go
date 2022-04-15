package main

import "fmt"

func main() {
	res := func() func() int {
		return func() int {
			i := 0
			i++
			return i
		}

	}
	fmt.Printf("%t", res)

}
