package main

import (
	"fmt"
	"golang-study-practices/words/server"
)

func main() {
	//excel.ImportData()
	/*db.RawQueryAllField()
	sourceFilePath := "/Users/CharlesDuan/Desktop/COCA20000增强版.xlsx"
	xlsx.OpenFile(sourceFilePath)*/
	var log = "Yes"
	fmt.Printf("application has been started. %s!!!", log)
	server.InitServer()
}
