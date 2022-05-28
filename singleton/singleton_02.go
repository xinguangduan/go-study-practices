// 饿汉模式，就是加载的时候全局初始化，缺点是启动占内存，慢
package main

import (
	"fmt"
)

//定义一个变量
var OneTerminal *terminalHungry

func init() {
	OneTerminal = newTerminalHungry()
}

//new函数，（比作：实例化类，new一个对象）
func newTerminalHungry() *terminalHungry {
	return &terminalHungry{}
}

//结构体，（比作：定义一个类）
type terminalHungry struct {
	//类型值（比作：类的某个属性）
}

//业务函数，（比作：类下的方法）
func (t *terminalHungry) Create(param string) (result interface{}, err error) {
	//处理param参数
	//业务逻辑处理
	//返回数据
	return
}

func main() {
	str := "helllo world"
	//引terminal 包时，已初始化变量Terminal，可直接通过调用（翻译翻译：引terminal 包时，1⃣已初始化实例类，可直接通过Terminal调用）
	var newTerminal = newTerminalHungry()
	result, err := newTerminal.Create(str)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	fmt.Println("hungry singleton create success", result)
}
