package jsonutils

import (
	"encoding/json"
	"github.com/Fcscanf/go-utils/httputils"
	"io"
	"net/http"
	"os"
)

// JsonUnmarshalByJsonFile 从JSON文件中反序列化JSON数据
func JsonUnmarshalByJsonFile(jsonPath string, v any) error {
	searchResultJson, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(searchResultJson, &v)
}

// JsonUnmarshalByURL 从URL请求中反序列化JSON数据
func JsonUnmarshalByURL(method, reqUrl, proxyUrl string, customHeader map[string]string, reqBody io.Reader, v any) error {
	resBody, err := httputils.GetResponseBodyFormUrl(method, reqUrl, proxyUrl, customHeader, reqBody)
	if err != nil {
		return err
	}
	return json.Unmarshal(resBody, &v)
}

// JsonUnmarshalByRequest 将Post请求的Body反序列化转为Struct
func JsonUnmarshalByRequest(request *http.Request, v any) error {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &v)
}
