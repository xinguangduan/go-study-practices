package main

import "fmt"

func main() {
	defer say1() // It's like pressing it into the stack. first in and last out
	defer say2()
	defer say3()

	fmt.Println("all print.")
}

func say1() {
	fmt.Println("say 1")
}
func say2() {
	fmt.Println("say 2")
}
func say3() {
	fmt.Println("say 3")
}
