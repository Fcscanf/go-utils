package httpinterceptor

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	BusinessError   = "System Error"
	BusinessSuccess = "Success"
)

type ResponseEntity struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r ResponseEntity) Ok(data any) ResponseEntity {
	r.Code = 200
	r.Message = BusinessSuccess
	r.Data = data
	return r
}

func (r ResponseEntity) Fail(data any) ResponseEntity {
	r.Code = 500
	r.Message = BusinessError
	r.Data = data
	return r
}

func (r ResponseEntity) FailMessage(msg string) ResponseEntity {
	r.Code = 500
	r.Message = msg
	return r
}

func (r ResponseEntity) Msg(code int, msg string, data any) ResponseEntity {
	r.Code = code
	r.Message = msg
	r.Data = data
	return r
}

func (r ResponseEntity) ResponseJson(writer http.ResponseWriter, result any, err error) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		log.Printf("servive error of %s", err)
		r.FailMessage(BusinessError)
	} else {
		r.Ok(result)
	}
	response, _ := json.Marshal(r)
	_, _ = writer.Write(response)
}
