package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("1000feng", "100"))
	fmt.Println(strings.ContainsRune("1000feng张三", '张'))
	fmt.Println(strings.Index("1000feng张三", "张三")) //8
	fmt.Println(strings.Count("1000feng张三", "张三")) //1
}
