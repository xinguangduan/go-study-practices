package main

import "fmt"

func main() {

	// add element by append,reduce element by [1:]
	var s1 = make([]string, 0, 5)
	tmp := append(s1, "a", "b", "c", "d", "e")
	fmt.Println(tmp)
	// break slice
	tmp01 := tmp[2:]
	fmt.Println(tmp01)
}
