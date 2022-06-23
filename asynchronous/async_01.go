package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
}

//一个查询结构体
type project struct {
	//参数Channel
	name   chan string
	result chan string
}

//addProject
func addProject(u user, p project) {
	fmt.Println("开始异步任务")

	//检查用户权限
	//checkPermission(u)

	//启动协程
	go func() {
		fmt.Println("另一个协程实行异步任务")
		time.Sleep(time.Second * 5)
		//获取输入
		name := <-p.name
		//访问数据库，输出结果通道
		p.result <- "add project :" + name
		fmt.Println("异步任务做完 耗时 5秒")
	}()

}

//主进程
func main() {
	//初始化project
	p := project{make(chan string, 1), make(chan string, 1)}
	//某位用户
	u := user{}
	//执行addProject，注意执行的时候还不需要告知要创建的项目名字
	addProject(u, p)

	//准备参数
	p.name <- "an-asynchronous-project" // 异步调用

	t := 1
	fmt.Println("正在干别的事。。。")
	// 可以在这里先干别的事
	time.Sleep(time.Second * time.Duration(t))
	fmt.Println("别的事干完。耗时", t, "秒")
	//获取结果
	fmt.Println(<-p.result) // 这里要等到异步任务做完，这里会阻塞主进程

	fmt.Println("done")
}
