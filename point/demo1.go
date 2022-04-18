package main

import "fmt"

type Student struct {
	name string
	sex  bool
	age  int8
	memo string
}

func main() {
	stud := Student{
		name: "zhangsan",
		sex:  false,
		age:  10,
		memo: `"zhang san is a good student"`,
	}

	var a *Student
	a = &stud
	fmt.Println(*a)
	fmt.Println(a.age, a.name, a.sex, a.memo)
	fmt.Println((*a).age, (*a).name, (*a).sex, (*a).memo)

}
