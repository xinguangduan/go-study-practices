package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func task() {
	fmt.Println("do something")
}
func main() {
	// 创建trace输出文件
	traceFile, err := os.Create("trace.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer traceFile.Close()

	//启动trace
	err = trace.Start(traceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	//正常业务逻辑
	go task()

	// 停止trace
	trace.Stop()
}
