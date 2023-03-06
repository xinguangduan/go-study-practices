package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 请求头中包含“Transfer-Encoding: chunked”表示支持HTTP分块传输
	req, err := http.NewRequest("GET", "http://localhost:8000/live", nil)
	if err != nil {
		fmt.Println("NewRequest error: ", err)
		return
	}
	req.Header.Set("Transfer-Encoding", "chunked")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do error: ", err)
		return
	}
	defer resp.Body.Close()

	// 根据Content-Length设置接收缓冲区大小
	buf := make([]byte, 4096)
	file, err := os.Create("output.flv")
	if err != nil {
		fmt.Println("Create error: ", err)
		return
	}
	defer file.Close()

	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Read error: ", err)
			return
		}
		if n == 0 {
			break
		}
		file.Write(buf[:n])
	}
}

func getRangeS(rangeHeader string, contentLength int64) (string, string) {
	const prefix, suffix = "bytes=", "-"
	rangeStr := rangeHeader[len(prefix):]
	rangeArr := strings.Split(rangeStr, suffix)
	start, _ := strconv.ParseInt(rangeArr[0], 10, 64)
	var end int64
	if len(rangeArr) > 1 && rangeArr[1] != "" {
		end, _ = strconv.ParseInt(rangeArr[1], 10, 64)
	}
	if end == 0 {
		end = contentLength - 1
	}
	return strconv.FormatInt(start, 10), strconv.FormatInt(end, 10)
}
