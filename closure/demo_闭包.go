package main

import "fmt"

func main() {
	res := Counter()
	for i := 0; i < 10; i++ {
		fmt.Printf("res:%v\n", res())
	}
	res2 := Counter()
	for i := 0; i < 10; i++ {
		fmt.Printf("res:%v\n", res2())
	}
}

func Counter() func() int {
	i := 0
	res := func() int {
		i++
		return i
	}
	fmt.Println("Counter:", res)
	return res
}
