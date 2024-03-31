package httputils

import "testing"

type User struct {
	Name string `url:"name"`
	Age  int    `url:"age"`
}

func TestQueryStringEncoder(t *testing.T) {
	encoder, err := QueryStringEncoder(User{
		Name: "Tom",
		Age:  12,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(encoder)
}
