package main

import "fmt"

// NoteBook 本质是一个指针
type NoteBook interface {
}

func showNotebook(book NoteBook) {
	fmt.Println(book)
	v, ok := book.(string)
	if !ok {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
}
func main() {
	showNotebook("ssss")
	showNotebook(2.31)
}
