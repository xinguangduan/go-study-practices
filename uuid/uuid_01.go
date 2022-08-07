package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	for i := 0; i < 100; i++ {
		GenUUID()
	}
}

func GenUUID() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	fmt.Printf("%s", out)
	return string(out)
}
