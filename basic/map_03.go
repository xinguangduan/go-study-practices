package main

import (
	"fmt"
	"strconv"
)

func main() {
	m1 := map[string]string{
		"a": "aaa-",
		"b": "bbb-",
		"c": "ccc-",
	}
	fmt.Println(m1)
	changeMapValue(m1)
	fmt.Println(m1)
}

// 传递的是map的指针，能修改其中的值
func changeMapValue(m map[string]string) {

	for i := 0; i < len(m); i++ {
		for k, v := range m {
			m[k] = v + strconv.Itoa(i)
		}
	}
}
