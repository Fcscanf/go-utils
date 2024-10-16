package httpclient

import "testing"

func TestGetFile(t *testing.T) {
	HttpClient{}.GetFile("https://pic2.zhimg.com/v2-5e8b41cae579722bd6b8a612bf1660e6.jpg", "D:\\test.jpg", nil)
}
