package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err) //todo
	}
	fmt.Printf("%s", out)
}
