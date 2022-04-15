package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi("222")
	fmt.Println(a + 2)

	b, _ := strconv.ParseBool("1")
	fmt.Println("result:", b)
}
