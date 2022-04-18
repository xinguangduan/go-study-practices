package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

//

var totalRecords int
var totalQuitCount int
var topicStatistics map[string]int
var domainStatistics map[string]int

// entities
type NLUResponseAnalysis struct {
	UserId    string `json:userid`
	ProductId string `json:productid`
	DeviceId  string `json:deviceid`
	Category  string `json:category`
	Domain    string `json:domain`
	Name      string `json:name`
	Query     string `json:query`
}

type NLUResponse struct {
	UserId    string `json:userid`
	ProductId string `json:productid`
	DeviceId  string `json:deviceid`
	Category  string `json:category`
	Payload   Payload
	Domain    string `json:domain`
	Name      string `json:name`
	Query     string `json:query`
}
type Payload struct {
	Result string `json:result`
}

type Intents struct {
	Intents []Intent `json:intents`
	Query   string   `json:query`
}

type Intent struct {
	Domain string `json:domain`
	Name   string `json:name`
}

func load(sourceFile string, analysisResultFile string) {
	fmt.Println("Source file is", sourceFile)
	fmt.Println("Analysis results file is ", analysisResultFile)

	file, _ := os.Open(sourceFile)
	if _, err2 := os.Stat(analysisResultFile); os.IsNotExist(err2) {
		os.Create(analysisResultFile)
	}
	res, err3 := os.OpenFile(analysisResultFile, os.O_RDWR|os.O_APPEND, 0666)
	if err3 != nil {
		fmt.Println(nil)
		return
	}

	defer file.Close()
	defer res.Close()
	stat, _ := file.Stat()
	fmt.Println("===analysis file size(kb):", stat.Size()/1024)
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = -1
	// init containers
	if topicStatistics == nil {
		topicStatistics = make(map[string]int)
	}

	nluWriter := bufio.NewWriter(res)
	for {
		row, err := reader.Read()
		if err != nil && err != io.EOF {
			fmt.Println(err)
		}
		if err == io.EOF {
			break
		}
		//fmt.Println(row)
		logType := row[14]
		totalRecords++

		val, ok := topicStatistics[logType]
		if ok {
			topicStatistics[logType] = val + 1
		} else {
			topicStatistics[logType] = 1
		}
		nluWriter.Flush()
		//fmt.Println("logType:",logType)
	}
	res.Sync()
}
func clearFileContent(fileName string) {
	os.Truncate(fileName, 0)
	fmt.Println("clear the file", fileName)
}
func main() {
	var basePath string
	var nluResults string
	flag.StringVar(&basePath, "basePath", "./", "")
	flag.StringVar(&nluResults, "nResults", "analysis.log", "nlu analysis results")

	//这里有一个非常重要的操作,转换， 必须调用该方法
	flag.Parse()
	fmt.Println("basePath=", basePath)
	fmt.Println("nluResults=", nluResults)
	fmt.Println("begin analysis ...")
	const sourceFile = "/Users/CharlesDuan/log/src/cui-tools-analysis/translation/data_20210625_062722.txt"
	const analysisResultFile = "/Users/CharlesDuan/log/src/cui-tools-analysis/translation/analysisResult.txt"
	nluResultsFilePath := sourceFile
	clearFileContent(nluResultsFilePath)

	load(sourceFile, analysisResultFile)
	fmt.Println("===total records==", totalRecords)
	for k, v := range topicStatistics {
		fmt.Println(k+":", v)
	}
}
