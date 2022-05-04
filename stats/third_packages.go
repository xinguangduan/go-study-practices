package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var packages = make(map[string]string)

func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "", "")
	flag.Parse()
	//打开目录
	f, err := os.OpenFile("/Users/charles/Desktop/poms", os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer f.Close()
	//读取目录项
	files, err := f.Readdir(-1) //-1 读取目录中的所有目录项
	if err != nil {
		fmt.Println("readdir err:", err)
		return
	}
	//变量返回切片
	for _, fileInfo := range files {
		if !fileInfo.IsDir() {
			if !strings.HasSuffix(fileInfo.Name(), ".xml") {
				fmt.Println("not support file:", fileInfo.Name())
				continue
			}
			read(filePath, fileInfo.Name())
		}
	}
	for k, _ := range packages {
		//jar := k + "-" + v + ".jar"
		fmt.Println(k)
	}
}

func read(filePath string, sourceFile string) {
	fmt.Println(sourceFile)
	file, _ := os.Open(filePath + sourceFile)
	defer file.Close()
	//scan line
	scanner := bufio.NewScanner(file)
	needFetch := false
	var artifactId, version string
	needAppendArtifactId := false
	needAppendVersion := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "<dependency>") {
			needFetch = true
		} else if strings.HasSuffix(line, "</dependency>") {
			needFetch = false
		}
		if !needFetch {
			continue
		}
		//fmt.Println(line)
		if needFetch && !strings.Contains(line, "dependency") {
			if strings.Contains(line, "artifactId") {
				artifactId = strings.Replace(line, "<artifactId>", "", -1)
				artifactId = strings.Replace(artifactId, "</artifactId>", "", -1)
				needAppendArtifactId = true
			}
			if strings.Contains(line, "version") {
				version = strings.Replace(line, "<version>", "", -1)
				version = strings.Replace(version, "</version>", "", -1)
				needAppendVersion = true
			}
			if strings.Contains(line, "optional") {
				version = "no"
				needAppendVersion = true
			}
			if strings.Contains(line, "exclusions") {
				version = "no"
				needAppendVersion = true
			}

			if needAppendArtifactId && needAppendVersion {
				packages[artifactId] = version
				//jar := artifactId + "-" + version + ".jar"
				//fmt.Println(jar)
				artifactId = ""
				version = ""
				needAppendVersion = false
				needAppendVersion = false

			}
		}
	}
}
