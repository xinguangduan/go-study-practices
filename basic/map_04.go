package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	mapInitTest01()
	mapInitTest02()
	mapInitTest03()
	mapInitTest04()
	mapInitTest05()
	mapInitTest06()
}

func mapInitTest01() {
	var domainMappingTree map[string]string
	if domainMappingTree == nil {
		fmt.Println("domainMappingTree is nil,going to make one")
		domainMappingTree = make(map[string]string)
	}
	domainMappingTree["slice"] = "12222"
	domainMappingTree["systemManager"] = "ssssss"
	fmt.Printf("%v \n", domainMappingTree)
}
func mapInitTest02() {
	domainMappingTree := make(map[string]string)
	domainMappingTree["slice"] = "12222"
	domainMappingTree["waste"] = "333333"
	domainMappingTree["systemManager"] = "ssssss"
	fmt.Printf("%v \n", domainMappingTree)
}
func mapInitTest03() {
	domainMappingTree := map[string]string{
		"slice":         "Test03_12222",
		"waste":         "Test03_333333",
		"systemManager": "Test03_ssssss",
	}
	var s = domainMappingTree["systemManager"]
	fmt.Printf("systemManger %s \n", s)
	fmt.Printf("%v \n", domainMappingTree)
}

func mapInitTest04() {
	domainMappingTree := map[string]string{
		"slice":         "Test03_12222",
		"waste":         "Test03_333333",
		"systemManager": "Test03_ssssss",
	}
	value, ok := domainMappingTree["systemManager"]

	if ok == true {
		fmt.Println("has been contained", value)
	} else {
		fmt.Println("has been contained")
	}
	fmt.Printf("%v \n", domainMappingTree)
}
func mapInitTest05() {
	domainMappingTree := map[string]string{
		"slice":         "Test03_12222",
		"waste":         "Test03_333333",
		"systemManager": "Test03_ssssss",
	}
	delete(domainMappingTree, "systemManager")
	fmt.Printf("%v \n", domainMappingTree)
	for k, v := range domainMappingTree {
		fmt.Println(k, v)
	}
	fmt.Println("count:" + strconv.Itoa(len(domainMappingTree)))
}
func mapInitTest06() {

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	// 修改
	ages["alice"] = 43
	fmt.Println(ages["alice"]) // 43

	// 删除
	delete(ages, "alice")

	// 因为我们一开始就知道names的最终大小，因此给slice分配一个合适的大小将会更有效。
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	events := map[string]string{}
	events["sarg"] = "sss"
	kv := events["sarg"]
	fmt.Println(kv)
}
