package main

import "fmt"

func main() {
	//string and byte convert each other
	s := "nihao zhangsan"
	b := []byte(s)
	fmt.Println(b)

	b1 := []byte{110, 105, 104, 97}

	s1 := string(b1)
	fmt.Println(s1)

}
