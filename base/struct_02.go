package main

import "fmt"

type Man struct {
	Name string
	Old  float64
}

func (m *Man) Eat() {
	fmt.Println("eat something", m.Name)
}

type SuperMan struct {
	Man
	High int
}

func (s SuperMan) Eat() {
	fmt.Println("sub class eat something", s.Name)
}

func main() {

	var s SuperMan

	s.Name = "zhangsan"
	s.Old = 30
	s.High = 199
	s.Eat()
	fmt.Println(s)

	s1 := SuperMan{Man{
		Name: "wangwu",
		Old:  19,
	}, 192}

	s1.Eat()
	fmt.Println(s1)
}
