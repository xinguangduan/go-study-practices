package main

import (
	"fmt"
)

func normalGoto() {
	var i = 0
AK:
	if i < 10 {
		fmt.Println(i)
		i++
		goto AK
	}
}

func main() {
	//normalGoto()
	//badGoto()
}

func badGoto() {
	var s = 0
Bad:
	if s < 10 {
		fmt.Println(s)
		goto Bad
	}

}
