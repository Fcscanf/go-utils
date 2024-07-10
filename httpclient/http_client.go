package httpclient

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
)

type HttpClient struct {
	http.Client
}

func (client HttpClient) SetProxyUrl(rawURL string) HttpClient {
	proxyUrl, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	t := &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Transport = t
	return client
}

// EnableCookieJar 开启HttpClient Cookie管理
func (client HttpClient) EnableCookieJar() HttpClient {
	// 创建一个新的cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("failed to create cookie jar: %v", err)
	}
	client.Jar = jar
	return client
}

func (client HttpClient) Get(url string, requestHeader map[string]string) (*http.Response, []byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	for key, value := range requestHeader {
		req.Header.Add(key, value)
	}
	return client.Execute(req)
}

func (client HttpClient) GetJson(url string, requestHeader map[string]string, v any) error {
	_, resBytes, err := client.Get(url, requestHeader)
	if err != nil {
		return err
	}
	return json.Unmarshal(resBytes, &v)
}

func (client HttpClient) Post(url string, reqBody io.Reader, requestHeader map[string]string) (*http.Response, []byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return nil, nil, err
	}
	for key, value := range requestHeader {
		req.Header.Add(key, value)
	}
	return client.Execute(req)
}

func (client HttpClient) PostJson(url string, reqBody io.Reader, requestHeader map[string]string, v any) error {
	_, resBytes, err := client.Post(url, reqBody, requestHeader)
	if err != nil {
		return err
	}
	return json.Unmarshal(resBytes, &v)
}

func (client HttpClient) PostForm(url string, formData, fileData, requestHeader map[string]string) (*http.Response, []byte, error) {
	uploadBody := &bytes.Buffer{}
	writer := multipart.NewWriter(uploadBody)
	for fileField, file := range fileData {
		f, err := os.Open(file)
		if err != nil {
			return nil, nil, err
		}
		fWriter, err := writer.CreateFormFile(fileField, filepath.Base(file))
		if err != nil {
			return nil, nil, err
		}

		_, err = io.Copy(fWriter, f)
		if err != nil {
			return nil, nil, err
		}
		err = f.Close()
		if err != nil {
			return nil, nil, err
		}
	}
	for k, v := range formData {
		_ = writer.WriteField(k, v)
	}
	err := writer.Close()
	if err != nil {
		return nil, nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, uploadBody)
	if err != nil {
		return nil, nil, err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	for key, value := range requestHeader {
		request.Header.Add(key, value)
	}
	return client.Execute(request)
}

func (client HttpClient) PostFormJson(url string, formData, fileData, requestHeader map[string]string, v any) error {
	_, resBytes, err := client.PostForm(url, fileData, fileData, requestHeader)
	if err != nil {
		return err
	}
	return json.Unmarshal(resBytes, &v)
}

func (client HttpClient) Execute(r *http.Request) (*http.Response, []byte, error) {
	res, err := client.Do(r)
	if err != nil {
		return res, nil, err
	}
	return after(res)
}

func after(response *http.Response) (*http.Response, []byte, error) {
	// 检查响应是否为 Gzip 压缩
	defer response.Body.Close()
	var result []byte
	var err error
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		// 创建 Gzip 读取器
		gzipReader, err := gzip.NewReader(response.Body)
		if err != nil {
			return response, nil, err // 如果创建失败，直接返回错误
		}
		defer gzipReader.Close()
		// 解压缩并读取响应数据
		result, err = io.ReadAll(gzipReader)
	default:
		// 非压缩响应数据的处理
		result, err = io.ReadAll(response.Body)
	}
	return response, result, err
}
