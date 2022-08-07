// 不用client的post请求
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	data := make(map[string]interface{})
	data["name"] = "zhaofan"
	data["age"] = "23"
	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
