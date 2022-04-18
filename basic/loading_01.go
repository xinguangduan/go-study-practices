package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		time.Sleep(50 * time.Millisecond)
		bar := strings.Repeat("=", i) + strings.Repeat(" ", 99-i)
		fmt.Printf("\r%.0f%%[%s]", float64(i)/99*100, bar)
	}
	fmt.Println()
	fmt.Println("all are loaded.")
}
