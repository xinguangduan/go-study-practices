package main

import (
	"fmt"
	"runtime"
)

func main() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.BySize)
	fmt.Printf("TotalAlloc: %v bytes\n", mem.TotalAlloc)
	fmt.Printf("Sys: %v bytes\n", mem.Sys)
}
