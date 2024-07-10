package jsonutils

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// JsonUnmarshalByJsonFile 从JSON文件中反序列化JSON数据
func JsonUnmarshalByJsonFile(jsonPath string, v any) error {
	jsonFileBytes, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonFileBytes, &v)
}

// JsonUnmarshalByRequest 将Post请求的Body反序列化转为Struct
func JsonUnmarshalByRequest(request *http.Request, v any) error {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &v)
}
