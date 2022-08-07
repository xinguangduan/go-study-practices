package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 把一些参数做成变量而不是直接放到url中
func main() {
	params := url.Values{}
	Url, err := url.Parse("http://httpbin.org/get")
	if err != nil {
		return
	}
	params.Set("name", "zhaofan")
	params.Set("age", "23")
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath) // https://httpbin.org/get?age=23&name=zhaofan
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
