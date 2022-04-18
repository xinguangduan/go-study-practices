package main

import (
	"fmt"
	"strings"
)

func main() {
	TestContain()
}
func TestContain() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}
