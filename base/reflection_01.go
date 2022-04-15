package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type MyBook struct {
	name string
}

func (b *MyBook) ReadBook() {
	fmt.Println("read book")
}

func (b *MyBook) WriteBook() {
	fmt.Println("write book")
}

func main() {
	var r Reader
	myBook := MyBook{"a book"}
	r = &myBook
	r.ReadBook()

	var w Writer
	w = r.(Writer) // pair<type:book>,value:book{}地址
	w.WriteBook()
}
