//GET请求添加请求头
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
	req.Header.Add("name", "zhaofan")
	req.Header.Add("age", "3")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}
