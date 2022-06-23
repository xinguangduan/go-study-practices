package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/* bufio.Reader和bufio.Scanner的关系

bufio.Reader是go早期的版本也是用来处理文本，使用起来有一些不方便，例如需要处理行太长的问题，而bufio.Scanner是go1.1中新增加的功能，既然是新加的功能肯定是修正之前的不足，在使用上更加方便，比如就不用处理行太长的问题。

总之就是bufio.Scanner是后开发的模块，功能更强大，使用更方便。
*/
func main() {
	loadFiles("/Users/charles/Desktop/attack/")
	fmt.Println("total:", index)
}

var index = 0

func loadFiles(fileDir string) {
	files, _ := ioutil.ReadDir(fileDir)
	for i, file := range files {
		log.Println(i, "====>read file", file.Name())
		err := handleText(fileDir + file.Name())
		if err != nil {
			//panic(err)
			return
		}
	}
}
func handleText(textFile string) error {
	file, err := os.Open(textFile)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", textFile, err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // or
		//line := scanner.Bytes()
		index++

		//do_your_function(line)
		fmt.Printf("%s\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Cannot scanner text file: %s, err: [%v]", textFile, err)
		return err
	}

	return nil
}
