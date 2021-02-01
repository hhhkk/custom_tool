package net

import (
	"fmt"
	"github.com/hhhkk/custom_tool/tool"
	"io"
	"net/http"
	"os"
	"strconv"
)

var packageSize int64 = 4096
var GOROUTINE_COUNT = 5

func SliceDownload(url, path string) bool {
	request := BuildRequest(url, "HEAD", nil)
	isSuccess := false
	Request(request, func(reader io.Reader, response *http.Response) {
		length := response.ContentLength
		if tool.IsExist(path) {
			size := tool.GetFileSize(path)
			if size == length {
				isSuccess = true
			} else {
				check(response, &isSuccess, url, path, length)
			}
		} else {
			check(response, &isSuccess, url, path, length)
		}
	}, nil)
	return isSuccess
}

func check(response *http.Response, isSuccess *bool, url string, path string, length int64) {
	tool.Open(path, func(file *os.File) {
		if response.Header.Get("Accept-Ranges") == "bytes" {
			startSliceDownload(url, file, length)
		} else {
			fmt.Println("DownloadByFile")
			*isSuccess = DownloadByFile(url, file)
		}
	}, nil)
}

func startSliceDownload(url string, file *os.File, length int64) {
	threadCount := int(length / packageSize)
	if length%packageSize > 0 {
		threadCount++
	}
	ints := make(chan int, GOROUTINE_COUNT)
	for i := 0; i < threadCount; i++ {
		start := int64(i) * packageSize
		end := start + packageSize
		if end > length {
			end = length
		}
		go slice(url, file, ints, start, end)
		ints <- 0
	}
}

func slice(url string, file *os.File, ints chan int, start int64, end int64) {
	request := BuildRequest(url, "GET", nil)
	request.Header.Set(
		"Range",
		"bytes="+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10))
	Request(request, func(reader io.Reader, response *http.Response) {
		if response.StatusCode == 206 {
			bytes := make([]byte, end-start)
			reader.Read(bytes)
			file.Write(bytes)
			file.Sync()
			fmt.Println("slice",start,end)
		}
	}, func(err error) {

	})
	<-ints
}
