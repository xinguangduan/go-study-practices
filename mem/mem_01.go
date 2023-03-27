package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前内存使用情况
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("当前内存使用情况：Alloc = %v MiB\n", m.Alloc/1024/1024)

	// 设置应用程序内存大小为 100 MiB
	//size := uint64(100 * 1024 * 1024)
	//runtime.SetMallocBytes(int64(size))
	//runtime.MemStats{}

	// 再次获取当前内存使用情况
	runtime.ReadMemStats(&m)
	fmt.Printf("当前内存使用情况：Alloc = %v MiB\n", m.Alloc/1024/1024)
}
