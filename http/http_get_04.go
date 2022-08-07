package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 解析返回类型为json的结果
type result struct {
	Args    string            `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

func main() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%#v", res)
}
