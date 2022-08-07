package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	urlValues := url.Values{
		"name": {"zhaofan"},
		"age":  {"23"},
	}
	reqBody := urlValues.Encode()
	resp, _ := http.Post("http://httpbin.org/post", "text/html", strings.NewReader(reqBody))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
