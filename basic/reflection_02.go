package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   string
	Name string
	Age  float64
}

func showUser(u interface{}) {
	t := reflect.TypeOf(u)
	fmt.Println(t)
	v := reflect.ValueOf(u)
	fmt.Println(v)
}

func main() {
	var u = User{"0001", "里斯", 19}
	showUser(u)
}
