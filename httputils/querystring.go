package httputils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func QueryStringDecoder4Request(r *http.Request, val any) error {
	t := reflect.TypeOf(val)
	if t.Kind() != reflect.Pointer {
		return fmt.Errorf("val must be of pointer type")
	}
	v := reflect.ValueOf(val).Elem()
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("url")
		if tag != "" && v.Field(i).Kind() == reflect.String {
			v.Field(i).SetString(r.FormValue(tag))
		}
	}
	return nil
}

func QueryStringEncoder(val any) (string, error) {
	if val == nil {
		return "", fmt.Errorf("val cannot be nil")
	}
	t := reflect.TypeOf(val)
	v := reflect.Value{}
	switch t.Kind() {
	case reflect.Struct:
		v = reflect.ValueOf(val)
	case reflect.Pointer:
		v = reflect.ValueOf(val).Elem()
	default:
		return "", fmt.Errorf("val must be of struct or pointer type")
	}
	result := "?"
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("url")
		if tag != "" && v.Field(i).Kind() == reflect.String {
			result += tag + "=" + v.Field(i).String() + "&"
		}
	}
	return strings.TrimSuffix(result, "&"), nil
}
