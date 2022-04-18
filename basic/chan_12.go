package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	resultCh := make(chan chan string, 5000)
	wg := sync.WaitGroup{}
	go replay(resultCh)
	startTime := time.Now()
	operation2(resultCh, "aaa", &wg)
	operation2(resultCh, "bbb", &wg)
	operation1(resultCh, "ccc", &wg)
	operation1(resultCh, "ddd", &wg)
	operation2(resultCh, "eee", &wg)
	wg.Wait()
	endTime := time.Now()
	fmt.Printf("Process time %s", endTime.Sub(startTime))
}

func replay(resultCh chan chan string) {
	for {
		//拿到一个chan 读取值 这个时候拿到的是先进先出 因为所有方法是按顺序加入chan的
		c := <-resultCh
		//读取嵌套chan中的值，这个时候等待3秒 因为是operation2中执行了3秒 在这3绵中 其实其余的4个方法也已经执行完毕。之后的方法则不需要等待sleep的时间
		r := <-c
		fmt.Println(r)
	}
}

func operation1(ch chan chan string, str string, wg *sync.WaitGroup) {
	//先创建一个chan 兵给到嵌套chan 占据一个通道 这个通道是阻塞的
	c := make(chan string)
	ch <- c
	wg.Add(1)
	go func(str string) {
		time.Sleep(time.Second * 1)
		c <- "operation1:" + str
		wg.Done()
	}(str)
}

func operation2(ch chan chan string, str string, wg *sync.WaitGroup) {
	c := make(chan string)
	ch <- c
	wg.Add(1)
	go func(str string) {
		time.Sleep(time.Second * 2)
		c <- "operation2:" + str
		wg.Done()
	}(str)
}
