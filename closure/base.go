package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	Relational()
	ScoreMark()
	Loop()
	a, b, c := Loop2()
	fmt.Println(a, b, c)
	s := GetStud()
	fmt.Printf("%T %v\n", s, s)
	fmt.Printf("%T %#v\n", s, s)
	fmt.Printf("%T %+v\n", s, s)
	std, _ := json.Marshal(s)
	fmt.Printf("%v\n", string(std))
	fmt.Printf("a+b=%d", sum(1, 4))

	var res bytes.Buffer
	json.Indent(&res, []byte(std), "", "    ")

	fmt.Printf("%s", res.String())

	test()
}

var lock sync.Mutex

func test() {
	lock.Lock()
	fmt.Printf("wwwwwww")
	lock.Unlock()
}
func testDefer() {
	lock.Lock()
	defer lock.Unlock()
}

type caseFunc func(a, b int) int

func sum(a, b int) int {
	return a + b
}

type Student struct {
	Id       string
	Name     string
	Password string
	Sex      int
}

func GetStud() (student Student) {
	student.Id = "1111"
	student.Name = "张三"
	student.Password = "20000"
	student.Sex = 1
	return
}

func Loop2() (a, b, c int) {
	a = 10
	b = 20
	c = 30
	return
}
func Loop() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("sub  ")
			if j == 4 {
				break
			}
		}
		fmt.Println("root  ")
	}
}

func ScoreMark() {
	score := 88

	if score >= 90 {
		fmt.Println("优秀")

	} else if score >= 80 {
		fmt.Println("良好")

	} else if score >= 70 {
		fmt.Println("中等")

	} else if score >= 60 {
		fmt.Println("一般")

	} else {
		fmt.Println("不及格")

	}

	switch score {
	case 80:
		fmt.Println("良好")
		fallthrough
	case 70:
		fmt.Println("中等")
	case 60:
		fmt.Println("一般")
	default:
		fmt.Println("不及格")
	}

}

func Relational() {
	a, b, c, d := 1, 2, "test", "test"

	if a == b {
		fmt.Printf("s")
	} else {
		fmt.Println("a")
	}

	if c == d {
		fmt.Printf("true")
	} else {
		fmt.Println("false")
	}
}
