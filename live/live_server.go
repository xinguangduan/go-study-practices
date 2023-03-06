package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/live", handleLiveStream)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleLiveStream(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取FLV文件路径
	filepath := filepath.Join(".", "live.flv")

	// 打开FLV文件
	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "video/x-flv")
	w.Header().Set("Content-Length", string(fileInfo.Size()))

	// 支持HTTP分块传输
	if rangeHeader := r.Header.Get("Range"); rangeHeader != "" {
		start, end := getRange(rangeHeader, fileInfo.Size())
		w.Header().Set("Content-Range", "bytes "+start+"-"+end+"/"+string(fileInfo.Size()))
		w.WriteHeader(http.StatusPartialContent)
		s, _ := strconv.ParseInt(start, 10, 64)
		_, err = file.Seek(s, 0)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		e, _ := strconv.ParseInt(end, 10, 64)
		io.CopyN(w, file, e-s+1)
	} else {
		w.WriteHeader(http.StatusOK)
		io.Copy(w, file)
	}
}

func getRange(rangeHeader string, contentLength int64) (string, string) {
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
