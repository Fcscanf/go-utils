package jsonutils

import (
	"encoding/json"
	"github.com/Fcscanf/go-utils/httputils"
	"io"
	"os"
)

// JsonUnmarshalByJsonFile 从JSON文件中反序列化JSON数据
func JsonUnmarshalByJsonFile(jsonPath string, v any) error {
	searchResultJson, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(searchResultJson, &v)
	if err != nil {
		return err
	}
	return err
}

// JsonUnmarshalByURL 从URL请求中反序列化JSON数据
func JsonUnmarshalByURL(method, reqUrl string, customHeader map[string]string, reqBody io.Reader, useProxy bool, v any) error {
	resBody, err := httputils.GetResponseBodyFormUrl(method, reqUrl, customHeader, reqBody, useProxy)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resBody, &v)
	if err != nil {
		return err
	}
	return err
}
