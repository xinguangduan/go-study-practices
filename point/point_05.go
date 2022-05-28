package main

import "fmt"

func main() {

	a := 10
	fmt.Printf("%x \n", &a)

	b := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%x \n", &b)

	var c = 100
	var ip *int

	ip = &c
	fmt.Println(ip)
	fmt.Println(*ip)
	fmt.Println(&*ip)
	fmt.Println(&(*ip))
	fmt.Println(*(&ip))

}
