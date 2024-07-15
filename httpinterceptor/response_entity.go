package httpinterceptor

const (
	BusinessError   = "System Error"
	BusinessSuccess = "Success"
)

type ResponseEntity struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r *ResponseEntity) Ok(data any) {
	r.Code = 200
	r.Message = BusinessSuccess
	r.Data = data
}

func (r *ResponseEntity) Fail(data any) {
	r.Code = 500
	r.Message = BusinessError
	r.Data = data
}

func (r *ResponseEntity) FailMessage(msg string) {
	r.Code = 500
	r.Message = msg
}

func (r *ResponseEntity) Msg(code int, msg string, data any) {
	r.Code = code
	r.Message = msg
	r.Data = data
}
