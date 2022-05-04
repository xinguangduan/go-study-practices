package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name   string   `json:"name,omitempty"`
	Price  float64  `json:"price,omitempty"`
	Age    float64  `json:"age,omitempty"`
	Actors []string `json:"actors"`
}

func main() {

	// Convert go struct to json string
	m := Movie{
		Name:   "大闹天宫",
		Price:  100,
		Age:    20,
		Actors: []string{"zhangsan", "lisi", "wangwu"},
	}
	jsonBytes, _ := json.Marshal(m)
	jsonStr := string(jsonBytes)
	fmt.Println(jsonStr)

	// Convert JSON to Go struct
	// {"name":"大闹天宫","price":100,"age":20,"actors":["zhangsan","lisi","wangwu"]}
	tmp := &Movie{}
	var sourceStr = "{\"name\":\"大闹天宫\",\"price\":100,\"age\":20,\"actors\":[\"zhangsan\",\"lisi\",\"wangwu\"]}"
	err := json.Unmarshal([]byte(sourceStr), tmp)
	if err != nil {
		fmt.Println("convert error")
	}

	fmt.Println(tmp)

}
