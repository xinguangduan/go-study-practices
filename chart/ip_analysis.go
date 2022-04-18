package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const SplitBeginTag = "peer:/"
const SplitEndTag = ":"
const AnalysisResults = "analysis_results.log"
const BasicPath = "log/"

var IPStatistics map[string]int

func analysisLogs(fileName string) {
	file, err := os.Open(BasicPath + fileName)
	if err != nil {
		fmt.Printf("%s", err)
		panic("load file error")
		return
	}
	defer file.Close()
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(file)
	if IPStatistics == nil {
		fmt.Println("items is nil,going to make one")
		IPStatistics = make(map[string]int)
	}
	var totalRecords int
	for {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		line := string(b)
		//fmt.Println(line)
		if !strings.Contains(line, SplitBeginTag) {
			continue
		}
		strs := strings.Split(line, SplitBeginTag)
		if len(strs) > 0 {
			substr := strs[1]
			inIP := strings.Split(substr, SplitEndTag)[0]
			strings.TrimSpace(inIP)
			count := 0
			if c, ok := IPStatistics[inIP]; ok {
				count = c + 1
			} else {
				count = 1
			}
			IPStatistics[inIP] = count
			//out := strconv.Itoa(count) + "=" + inIP
			//fmt.Println(out)
			totalRecords++
		}
	}
	summary := "==== file " + fileName + " total request records:" + strconv.Itoa(totalRecords) + ",statics map ip records:" + strconv.Itoa(len(IPStatistics))
	fmt.Println(summary)
	writeAnalysis(AnalysisResults, summary)

}

func writeAnalysis(filename, content string) error {
	// 写入文件
	// 判断文件是否存在
	filepath := BasicPath + filename
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		os.Create(filepath)
	}
	fd, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	defer fd.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(fd)
	_, err2 := w.WriteString(content + "\n")
	if err2 != nil {
		return err2
	}
	w.Flush()
	fd.Sync()
	return nil
}
func generateAnalysisResult() {
	for k, v := range IPStatistics {
		out := strconv.Itoa(v) + "=" + k
		//fmt.Println(out)
		writeAnalysis(AnalysisResults, out)
	}
}

func main() {
	//analysisLogs("req1.log")
	//analysisLogs("req2.log")
	//analysisLogs("req3.log")
	//analysisLogs("req4.log")
	//generateAnalysisResult()

	fmt.Printf("%5f", 1222)
}
