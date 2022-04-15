package main

import (
	"fmt"
	"strings"
)

func main() {
	// 一个汉字占3个字节
	s := "我爱GO语音"
	fmt.Println(len(s))
	for i, ch := range []rune(s) {
		fmt.Printf("%d,%c \n", i, ch)
	}
	has := strings.Contains(s, "G")
	fmt.Printf("%t", has)
}
