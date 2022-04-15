package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToLower("a A"))
	fmt.Println(strings.ToUpper("a A"))

	fmt.Println(strings.Compare("a", "b"))

	fmt.Println(strings.Replace("张三丰你好啊,张三丰你好啊", "张三", "李四", 1))
	s := []string{"aaa", "bbbb", "ccc"}
	d := []string{"ddd", "eeee", "gggg"}
	x := strings.Join(s, ",") + "," + strings.Join(d, ",")
	fmt.Println(strings.Join(s, ","))
	fmt.Println(strings.Join(d, ","))
	fmt.Println(x)
}
