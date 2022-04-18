package main

import "fmt"

func func1(name string, pw string) string {
	return name + pw
}
func func2(name string, password string) (mix string) {
	mix = name + password
	return
}

func func3(flag string, name ...string) string {
	names := flag
	for _, v := range name {
		names = names + v
	}
	return names
}

func func4(num ...int) (sum int, max int, min int) {

	for _, v := range num {
		if v > max {
			max = v
		}
		if v <= min {
			min = v
		}
		sum = sum + v
	}
	return
}

func main() {
	name := "zhangsan"
	password := "122345"
	mix := func1(name, password)
	fmt.Println(mix)
	mix = func2(name, password)
	fmt.Println(mix)
	flag := "flag"
	names := func3(flag, name, name)
	fmt.Println(names)
	sum, max, min := func4(1, 2, 3, 4)
	fmt.Println(sum, max, min)
}
