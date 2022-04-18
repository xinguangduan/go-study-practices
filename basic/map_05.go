package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter = struct {
		sync.RWMutex //读写锁
		m            map[string]int
	}{m: make(map[string]int)}

	counter.RLock() // 读锁定
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println(n)

	counter.Lock() // 写锁定
	counter.m["some_key"] = 1
	n = counter.m["some_key"]
	counter.Unlock()

	fmt.Println(n)
}
