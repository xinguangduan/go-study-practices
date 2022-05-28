package main

import "fmt"

type Ball struct {
	name string
	size int
}

var ballMap map[string]Ball

func init() {
	ballMap = make(map[string]Ball)
}

func addBall(balls map[string]Ball) {
	balls["www"] = Ball{
		name: "wangwu",
		size: 111110,
	}
}

func main() {

	ballMap["b1"] = Ball{
		name: "zhangsan",
		size: 110,
	}
	ballMap["22"] = Ball{
		name: "lisa",
		size: 200,
	}
	for _, ball := range ballMap {
		fmt.Println(ball)
	}
	fmt.Println("------------------")
	addBall(ballMap)

	for _, ball := range ballMap {
		fmt.Println(ball)
	}
}
