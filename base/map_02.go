package main

import "fmt"

func main() {
	//var m1 map[string]string
	//m1 = make(map[string]string)

	m1 := make(map[string]string)
	m1["a"] = "hello"
	m1["b"] = "nihao"
	m1["c"] = "bye"

	for k, v := range m1 {
		fmt.Println(k, v)
	}
	m1["d"] = "add new value"
	outputMap(m1)

	delete(m1, "d")
	outputMap(m1)

	m1["d"] = "modify old value"
	outputMap(m1)
	k := "a"
	v, ok := m1[k]
	if ok {
		fmt.Println("key d value exists", v)
	}

}

func outputMap(m map[string]string) {

	if len(m) <= 0 {
		return
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
