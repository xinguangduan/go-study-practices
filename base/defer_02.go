package main

import "fmt"

func main() {
	s := handle("tom")
	fmt.Println(s)
}

func handle(str string) string {
	defer closeSession()
	return "hello " + str
}
func closeSession() {
	fmt.Println("close the session")
}
