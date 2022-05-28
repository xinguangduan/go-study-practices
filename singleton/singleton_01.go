// 懒汉模式，加锁，要用的时候初始化
package main

import (
	"fmt"
	"sync"
)

//定义一个变量
var newTerminalLazy *terminal

var once sync.Once

// NewTerminalLazy new函数，（比作：实例化类，new一个对象）
func NewTerminalLazy() *terminal {
	//只执行一次，sync.Once方法已实现了双层锁的功能
	once.Do(func() {
		newTerminalLazy = new(terminal)
	})
	return newTerminalLazy
}

//结构体，（比作：定义一个类）
type terminal struct {
	//类型值（比作：类的某个属性）
}

//业务函数，（比作：类下的方法）
func (t *terminal) Create(param string) (result interface{}, err error) {
	//处理param参数
	//业务逻辑处理
	//返回数据
	return
}

func main() {
	str := "helllo world"
	//引terminal 包时，已初始化变量Terminal，可直接通过调用（翻译翻译：引terminal 包时，1⃣已初始化实例类，可直接通过Terminal调用）
	var newTerminal = NewTerminalLazy()
	result, err := newTerminal.Create(str)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	fmt.Println("lazy singleton create success", result)
}
