package main

import "fmt"

func main() {
	m := make(map[string]string)
	fmt.Println(m)
	m["sss"] = "bbb"
	fmt.Println(m)

	m1 := make(map[int]int)
	m1[1] = 1
	fmt.Println(m1)

	var m2 map[string]string
	// 使用之前需要make命令分配空间
	m2 = make(map[string]string)
	m2["a"] = "s"
	m2["b"] = "d"
	fmt.Println(m2)

	m3 := map[string]string{
		"k1": "abc",
		"k2": "cdb",
		"k3": "sss",
	}
	fmt.Println(m3)
}
