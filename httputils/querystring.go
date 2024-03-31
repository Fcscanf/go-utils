package httputils

import (
	"fmt"
	"net/http"
	"reflect"
)

func QueryStringDecoder4Request(r *http.Request, val any) error {
	t := reflect.TypeOf(val)
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("unsupported argument type")
	}
	v := reflect.ValueOf(val)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("url")
		if tag != "" && v.Field(i).Kind() == reflect.String {
			v.Field(i).SetString(r.FormValue(tag))
		}
	}
	return nil
}
