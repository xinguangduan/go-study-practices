package main

import "fmt"

func showBookType(a interface{}) {
	fmt.Println("sss", a)
}

type BookShop struct {
	name    string
	address string
}

func main() {

	var bs BookShop
	bs.address = "beijing"
	bs.name = "xinhuabookshop"
	fmt.Println(bs)
	showBookType(bs)

	book := &BookShop{
		name:    "zhangsan",
		address: "beijing",
	}
	fmt.Println(*book)

	bs1 := BookShop{"beijing", "ssss"}
	showBookType(bs1)

	showBookType("sss")

	// 通过接口获取具体类型
	var allType interface{}
	a := "abc"
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)

}
