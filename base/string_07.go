package main

import (
	"fmt"
	"strings"
)

func main() {
	TestFieldFunc()
}

func TestFieldFunc() {
	s := " 111, 2222, 344 ,XXS ,sss "
	fmt.Println(strings.Fields(s))
	fmt.Println(strings.SplitAfter(s, ","))
	fmt.Println("%Q", strings.Split(s, ","))
	cleanStr := strings.Trim(s, " ")
	splitStr := strings.SplitAfterN(cleanStr, ",", 2)

	for s, v := range splitStr {
		fmt.Println(s, strings.TrimSpace(v))
	}

	fmt.Println(strings.SplitAfter(cleanStr, ","))
	fmt.Println(strings.SplitAfter(cleanStr, ","))

}
