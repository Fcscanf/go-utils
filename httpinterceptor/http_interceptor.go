package httpinterceptor

import (
	"encoding/json"
	"github.com/Fcscanf/go-utils/httputils"
	"github.com/Fcscanf/go-utils/jsonutils"
	"log"
	"net/http"
)

var DefaultResponseJson = func(writer http.ResponseWriter, result any, err error) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r := &ResponseEntity{}
	if err != nil {
		log.Printf("servive error of %s", err)
		r.FailMessage(BusinessError)
	} else {
		r.Ok(result)
	}
	response, _ := json.Marshal(r)
	_, _ = writer.Write(response)
}

type serviceFunc[T, R any] func(T) (R, error)

func PostJsonService[T, R any](writer http.ResponseWriter, request *http.Request, v T, service serviceFunc[T, R]) {
	err := jsonutils.JsonUnmarshalByRequest(request, &v)
	if err != nil {
		DefaultResponseJson(writer, nil, err)
		return
	}
	result, err := service(v)
	DefaultResponseJson(writer, result, err)
}

func GetService[T, R any](writer http.ResponseWriter, request *http.Request, v T, service serviceFunc[T, R]) {
	err := httputils.QueryStringDecoder4Request(request, &v)
	if err != nil {
		DefaultResponseJson(writer, nil, err)
		return
	}
	result, err := service(v)
	DefaultResponseJson(writer, result, err)
}

type voidServiceFunc[T any] func(T) error

func PostJsonVoidService[T any](writer http.ResponseWriter, request *http.Request, v T, service voidServiceFunc[T]) {
	if err := jsonutils.JsonUnmarshalByRequest(request, &v); err != nil {
		DefaultResponseJson(writer, nil, err)
		return
	}
	DefaultResponseJson(writer, nil, service(v))
}

func GetVoidService[T any](writer http.ResponseWriter, request *http.Request, v T, service voidServiceFunc[T]) {
	if err := httputils.QueryStringDecoder4Request(request, &v); err != nil {
		DefaultResponseJson(writer, nil, err)
		return
	}
	DefaultResponseJson(writer, nil, service(v))
}
