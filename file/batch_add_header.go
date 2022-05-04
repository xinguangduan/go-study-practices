package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var sg sync.WaitGroup

func main() {
	var sourcePath, header, extends, mode string
	flag.StringVar(&sourcePath, "s", "/Users/charles/Desktop/未命名文件夹/src", "File path to be modified")
	flag.StringVar(&header, "c", "// Copyright (c) 2019-present Lenovo.  All rights reserved\n// Confidential and Proprietary\n", "The content what need to insert file header")
	flag.StringVar(&extends, "extends", ".java,", "File name extension,Multiple use `,` separate,for example:.java,.xml")
	flag.StringVar(&mode, "mode", "n", "File name extension,Multiple use `,` separate,for example:.java,.xml")

	flag.Parse()

	fmt.Println("source file path:", sourcePath)
	fmt.Println("header content:", header)
	fmt.Println("file extends:", extends)
	fmt.Println("replace mode:", extends)

	ts := time.Now()
	var files []string
	var fileExtensions = strings.Split(extends, ",")

	root := sourcePath
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		for _, extension := range fileExtensions {
			if extension == "" {
				continue
			}
			if strings.HasSuffix(path, extension) {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Println(file)
		sg.Add(1)
		go insertHeader(file, header)
	}
	sg.Wait()
	duration := time.Now().Sub(ts)
	fmt.Println("cost", duration)
}

/**
实现方法:
要知道插入的内容需要插入至哪个位置，先得定位到这个位置，可以采用逐行读取、匹配(可用正则匹配)，再重新写入的方式
实现步骤:
1.使用ioutil.ReadFile读取文件
2.将文件字节数组转换为字符串并按\n转化成切片，即每一行
3.遍历每一行内容，并将内容复制到另一个新的切片中，在正则匹配到符合条件的行后，将要插入到内容插入至新的切片中
4.将新切片内容转换成字符串后输入至文件中，替换原文件内容，即实现了在原文件指定行后插入内容
*/
var header1 = "// Copyright (c) 2019-present Lenovo.  All rights reserved"
var header2 = "// Confidential and Proprietary"

func insertHeader(fileName, content string) error {
	defer sg.Done()
	lineBytes, err := ioutil.ReadFile(fileName)
	var lines []string
	if err != nil {
		fmt.Println(err)
	} else {
		contents := string(lineBytes)
		lines = strings.Split(contents, "\n")
	}
	var newLines []string
	newLines = append(newLines, content)
	isContainHeader1 := false
	isContainHeader2 := false
	blankLine := 0
	for _, line := range lines {
		// replace
		if !isContainHeader1 {
			isContainHeader1 = strings.HasPrefix(line, header1)
			line = strings.ReplaceAll(line, header1, "")
		}
		if !isContainHeader2 {
			isContainHeader2 = strings.HasPrefix(line, header2)
			line = strings.ReplaceAll(line, header2, "")
		}
		if isContainHeader1 && isContainHeader2 {
			continue
		}
		if line == "" {
			if blankLine < 2 {
				blankLine = blankLine + 1
				continue
			}
		}
		newLines = append(newLines, line)
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	defer file.Close()
	_, err = file.WriteString(strings.Join(newLines, "\n"))
	if err != nil {
		return err
	}
	return nil
}

func appendStringInFile(filePath, content string) {
	defer sg.Done()
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("文件打开失败: %v", err)
	}
	defer file.Close()
	// 查找文件末尾的偏移量
	//n, _ := file.Seek(0, io.SeekStart)
	// 从末尾的偏移量开始写入内容
	_, err = file.WriteString("\n" + content + "\n")
	if err != nil {
		log.Fatalf("文件写入失败: %v", err)
	}
}
