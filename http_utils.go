package utils

import (
	"compress/gzip"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func HttpProxyClient() *http.Client {
	proxyUrl, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		panic(err)
	}
	t := &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: t,
		Timeout:   time.Duration(10 * time.Second),
	}
	return &client
}

func GetResponseBodyFormUrl(method, reqUrl string, customHeader map[string]string, reqBody io.Reader, useProxy bool) ([]byte, error) {
	request, _ := http.NewRequest(method,
		reqUrl, reqBody)
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	for key, value := range customHeader {
		request.Header.Add(key, value)
	}
	httpClient := http.DefaultClient
	if useProxy {
		httpClient = HttpProxyClient()
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	// 检查响应是否为 Gzip 压缩
	if response.Header.Get("Content-Encoding") == "gzip" {
		// 创建 Gzip 读取器
		gzipReader, err := gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
		defer gzipReader.Close()
		// 解压缩并读取响应数据
		return io.ReadAll(gzipReader)
	} else {
		// 非压缩响应数据的处理
		return io.ReadAll(response.Body)

	}
}
