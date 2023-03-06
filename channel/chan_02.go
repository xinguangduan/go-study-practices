package main

import (
	"fmt"
	"time"
)

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel 是 nil 的, 不能使用，需要先创建通道.")
		a = make(chan int, 1)
		fmt.Printf("数据类型是：%T,%p\n", a, a)
	}
	//a <- 1
	fmt.Printf("数据类型是：%T,%p\n", a, a)
	//send(a)
	//recv(a)
	sendAndRecv(a)

}
func test01(t chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中，i：", i)
			t <- i
		}
		// 循环结束后，向通道中写数据，表示要结束了。。

		fmt.Println("结束。。")

	}()

	data := <-t // 从ch1通道中读取数据
	fmt.Println("data-->", data)
	fmt.Println("main。。over。。。。")

	fmt.Printf("===%T,%p\n", t, t)

	//send(t)
	//recv(t)
	sendAndRecv(t)
}

func send(t chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中，i：", i)
			t <- i
		}
		// 循环结束后，向通道中写数据，表示要结束了。。
		fmt.Println("结束。。")
	}()
}

func recv(t chan int) {
	go func() {
		v := <-t
		fmt.Println("recv->", v)
	}()
}
func sendAndRecv(t chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中，i：", i)
			t <- i
			v := <-t
			fmt.Println("recv->", v)
		}
		// 循环结束后，向通道中写数据，表示要结束了。。
		fmt.Println("结束。。")
	}()

	//go func() {
	//	v := <-t
	//	fmt.Println("recv->", v)
	//}()
	time.Sleep(time.Second * 10)
}
