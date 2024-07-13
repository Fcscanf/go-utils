package httpinterceptor

import (
	"github.com/Fcscanf/go-utils/httputils"
	"github.com/Fcscanf/go-utils/jsonutils"
	"net/http"
)

type HttpResponse interface {
	ResponseJson(writer http.ResponseWriter, result any, err error)
}

var DefaultHttpResponse ResponseEntity

type serviceFunc[T, R any] func(T) (R, error)

func PostJsonService[T, R any](writer http.ResponseWriter, request *http.Request, v T, service serviceFunc[T, R]) {
	err := jsonutils.JsonUnmarshalByRequest(request, &v)
	if err != nil {
		DefaultHttpResponse.ResponseJson(writer, nil, err)
		return
	}
	result, err := service(v)
	DefaultHttpResponse.ResponseJson(writer, result, err)
}

func GetService[T, R any](writer http.ResponseWriter, request *http.Request, v T, service serviceFunc[T, R]) {
	err := httputils.QueryStringDecoder4Request(request, &v)
	if err != nil {
		DefaultHttpResponse.ResponseJson(writer, nil, err)
		return
	}
	result, err := service(v)
	DefaultHttpResponse.ResponseJson(writer, result, err)
}

type voidServiceFunc[T any] func(T) error

func PostJsonVoidService[T any](writer http.ResponseWriter, request *http.Request, v T, service voidServiceFunc[T]) {
	err := jsonutils.JsonUnmarshalByRequest(request, &v)
	if err != nil {
		DefaultHttpResponse.ResponseJson(writer, nil, err)
		return
	}
	DefaultHttpResponse.ResponseJson(writer, nil, service(v))
}

func GetVoidService[T any](writer http.ResponseWriter, request *http.Request, v T, service voidServiceFunc[T]) {
	err := httputils.QueryStringDecoder4Request(request, &v)
	if err != nil {
		DefaultHttpResponse.ResponseJson(writer, nil, err)
		return
	}
	DefaultHttpResponse.ResponseJson(writer, nil, service(v))
}
