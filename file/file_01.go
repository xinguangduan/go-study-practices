package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	file, err := os.Open(exPath + "/" + "ref_01.go")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	defer file.Close()
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(file)

	fmt.Println(buf)
}
