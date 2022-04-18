package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 200; i++ {
		time.Sleep(30 * time.Millisecond)
		bar := strings.Repeat("@", i) + strings.Repeat(" ", 199-i)
		fmt.Printf("\r%.0f%%[%s]", float64(i)/199*100, bar)
		os.Stdout.Sync()
	}
	fmt.Println("\n all data are loaded.")
}
