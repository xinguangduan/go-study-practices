package main

import "fmt"

// Animal 定义一个接口
type Animal interface {
	jump()
	getColor() string
}

// Cat 定义实现类，需要实现接口的所有方法
type Cat struct {
}

func (c *Cat) jump() {
	fmt.Println("do jump...")
}

func (c *Cat) getColor() string {
	fmt.Println("get color...")
	return "black"
}

func main() {
	animal := new(Animal)
	fmt.Println(animal)
	// 定义接口类型
	var a Animal

	a = &Cat{}
	a.jump()
	color := a.getColor()
	fmt.Println(color)

	var cat Cat
	cat.jump()
	cat.getColor()
	fmt.Println(cat)
}
