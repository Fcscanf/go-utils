package httputils

import (
	"fmt"
	"net/http"
	"reflect"
)

func QueryStringDecoder4Request(r *http.Request, val any) error {
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val).Elem()
	if t.Kind() != reflect.Pointer || v.IsNil() {
		return fmt.Errorf("unsupported argument type")
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("url")
		if tag != "" && v.Field(i).Kind() == reflect.String {
			v.Field(i).SetString(r.FormValue(tag))
		}
	}
	return nil
}
