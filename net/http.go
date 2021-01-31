package net

import (
	"compress/gzip"
	"crypto/tls"
	"github.com/hhhkk/custom_tool/log"
	"github.com/hhhkk/custom_tool/tool"
	"github.com/andybalholm/brotli"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var defaultClient *http.Client

var hookClient *http.Client

var hookUseClient func() bool

var hookRequestUA func() string

func init() {
	transport := http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		IdleConnTimeout:       time.Minute * time.Duration(3),
		TLSHandshakeTimeout:   time.Minute * time.Duration(3),
		ResponseHeaderTimeout: time.Minute * time.Duration(3),
		ExpectContinueTimeout: time.Minute * time.Duration(3),
	}
	defaultClient = &http.Client{
		Transport: &transport,
	}
}

func HookRequestUA(obj func() string) {
	hookRequestUA = obj
}

func HookUseClient(obj func() bool) {
	hookUseClient = obj
}

func HookHttpTransport(obj *http.Transport) {
	hookClient = &http.Client{
		Transport: obj,
	}
}

func BuildRequest(url, method string, body io.Reader) *http.Request {
	if request, err := http.NewRequest(method, url, body); err == nil {
		if hookRequestUA != nil {
			ua := hookRequestUA()
			request.Header.Add("User-Agent", ua)
		}
		return request
	} else {
		log.Fatal(err)
	}
	return nil
}

func BuildRequestEnableGzip(url, method string, body io.Reader) *http.Request {
	request := BuildRequest(url, method, body)
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	return request
}

func GetNetworkFileSize(url string) int64 {
	request := BuildRequest(url, "HEAD", nil)
	if resp, err := getClient().Do(request); err == nil {
		defer resp.Body.Close()
		return resp.ContentLength
	} else {
		return -1
	}
}

func BuildProxy(ip, port string) *func(*http.Request) (*url.URL, error) {
	if porxyUrl, err := url.Parse(ip + port); err == nil {
		obj := http.ProxyURL(porxyUrl)
		return &obj
	} else {
		log.Fatal(err)
	}
	return nil
}

func Download(url, path string) bool {
	isSuccess := false
	if tool.IsExist(path) {
		size := tool.GetFileSize(path)
		fileSize := GetNetworkFileSize(url)
		if size != fileSize {
			tool.Open(path, func(file *os.File) {
				isSuccess = DownloadByFile(url, file)
			}, nil)
		}
	} else {
		tool.Open(path, func(file *os.File) {
			isSuccess = DownloadByFile(url, file)
		}, nil)
	}
	return isSuccess
}

func DownloadByFile(url string, file *os.File) bool {
	request := BuildRequest(url, "GET", nil)
	Success := false
	Request(request, func(reader io.Reader, resp *http.Response) {
		if _, err := io.Copy(file, reader); err == nil {
			Success = true
		}
	}, func(err error) {
		Success = false
	})
	return Success
}

func getClient() *http.Client {
	if hookUseClient != nil && hookUseClient() {
		return hookClient
	} else {
		return defaultClient
	}
}

func Request(req *http.Request, success func(io.Reader, *http.Response), fail func(error)) {
	client := getClient()
	if resp, err := client.Do(req); err == nil {
		defer resp.Body.Close()
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			if reader, err := gzip.NewReader(resp.Body); err == nil {
				if success != nil {
					success(reader, resp)
				}
			} else {
				log.E(err)
				if fail != nil {
					fail(err)
				}
			}
		case "br":
			reader := brotli.NewReader(resp.Body)
			if success != nil {
				success(reader, resp)
			}
		default:
			if success != nil {
				success(resp.Body, resp)
			}
		}
	} else {
		log.E(err)
		if fail != nil {
			fail(err)
		}
	}
}
