package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type Student struct {
	Name   string   `json:"name,omitempty"`
	Price  int      `json:"price,omitempty"`
	Age    int      `json:"age,omitempty"`
	Actors []string `json:"actors,omitempty"`
}

//DO NOT EDIT
func main() {
	ts := time.Now()
	var obj = "{\"name\":\"大闹天宫\",\"price\":100,\"age\":20,\"actors\":[\"zhangsan\",\"lisi\",\"wangwu\"]}"
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	s := Student{}
	json.Unmarshal([]byte(obj), &s)
	fmt.Println(s)

	s.Age = 500
	s.Name = "美猴王"
	out, _ := json.Marshal(&s)
	fmt.Println(string(out))
	fmt.Println(time.Now().Sub(ts).Microseconds())
}
