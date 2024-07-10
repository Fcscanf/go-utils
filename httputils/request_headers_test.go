package httputils

import (
	"log"
	"testing"
)

func TestRequestHeaders(t *testing.T) {
	requestHeader := RequestHeader{}.
		ContentType("application/json; charset=utf-8").
		Referer("http://www.baidu.com").
		Add("k", "v")
	log.Println(requestHeader)
}
