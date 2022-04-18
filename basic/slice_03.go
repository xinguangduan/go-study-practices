package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 声明slice是一个切片，并且初始化，默认值1，2，3，长度是3
	s := []int{1, 2, 3}
	fmt.Println(s)
	// 声明slice是个切片，但没有给它分配空间
	var s1 []int
	fmt.Println(s1)
	//  声明slice，并给slice分配空间，初始化

	s2 := make([]int, 3)
	fmt.Println(s2)
	s3 := append(s2, 122)
	fmt.Println(s3)

	s4 := append(s3, 2, 3, 2)
	fmt.Println(s4)

	if s == nil {
		fmt.Println("s is a empty slice")
	} else {
		fmt.Println("s is not a empty slice")
	}

	s4 = make([]int, 0)
	if len(s4) == 0 {
		fmt.Println("s4 is a empty slice", cap(s4))
	} else {
		fmt.Println("s4 is not a empty slice", cap(s4))
	}

	var s5 = make([]string, 3, 10)
	fmt.Println(s5, cap(s5))

	// bad case, index out of range
	for i := 0; i < 11; i++ {
		s5[i] = strconv.Itoa(i)
	}
}
