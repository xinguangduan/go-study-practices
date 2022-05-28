package main

import "fmt"

type Person struct {
	name string
	age  int
	car  Car
}
type Car struct {
	name string
}

var personMap map[string]Person
var note = `
所以得出结论，当我们需要修改结构体的变量内容的时候，方法传入的结构体变量参数需要使用指针，也就是结构体的地址

需要修改map中的架构体的变量的时候也需要使用结构体地址作为map的value

如果仅仅是读取结构体变量，可以不使用指针，直接传递引用即可

*type 这里的type这个变量存放的东西是地址，这点需要明确，需要使用&type获取到地址。
`

func main() {
	pt := fmt.Println
	c := Car{
		"sss",
	}
	p := Person{
		name: "张三",
		age:  20,
		car:  Car{"宝马"},
	}
	pt(p)
	setCarName(&c, "奔驰")
	pt(c)
	personMap = make(map[string]Person)
	personMap["p1"] = p
	for _, person := range personMap {
		pt(person)
	}

}
func setCarName(car *Car, name string) {
	car.name = name
}
func (p *Person) setName(name string) {
	p.name = name
}
