package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONBody http body结构体
type JSONBody struct {
	StatusCode string      `json:"code"`
	Message    string      `json:"message"`
	Body       interface{} `json:"body,omitempty"`
}

var code = 0     /*default code*/
var message = "" /*default message*/

// JComplete 将结果反序列化为JSON，返回给客户端
// status_code : http status code
// code : custom code
// message : client message or http status message or err msg
// body : custom <struct body> message
var JComplete = func(w http.ResponseWriter, status_code int, code int, message string, body interface{}) {
	var msg = message
	if msg == "" {
		msg = http.StatusText(status_code)
	}

	var res = &JSONBody{
		StatusCode: fmt.Sprintf("%d.%d", status_code, code),
		Message:    msg,
		Body:       body,
	}

	/*同时设置header和http code，先设置header，然后code，否则header不生效*/
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status_code)

	var buffer = bytes.NewBufferString("")
	var enc = json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "    ")
	var _ = /*error ignore*/ enc.Encode(res)

	//var b, _ = json.Marshal(res)
	fmt.Fprint(w, buffer.String())
	return
}

// DecoratorFunc 包装路由器
// @f 路由函数
// @funcs 验证函数
func DecoratorFunc(f func(http.ResponseWriter, *http.Request), funcs ...func(*http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer CatchError(w, r)

		for _, fn := range funcs {
			var err = fn(r)
			if err != nil {
				JReject(w, http.StatusBadRequest, code, err.Error())
				return
			}
		}

		f(w, r)
		return
	}
}

// CatchError 用于路由上最终的panic捕获
var CatchError = func(w http.ResponseWriter, r *http.Request) {
	var v = recover()
	if v != nil {
		JReject(w, http.StatusInternalServerError, code, fmt.Sprint(v))
	}
}

// JRouteReject for rest json
type JRouteReject struct {
	StatusCode int    /*http code*/
	StatusText string /*http message*/
}

func (rr *JRouteReject) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	JReject(w, rr.StatusCode, code, rr.StatusText)
}

// NotFound 将结果定义为json格式
var NotFound = JRouteReject{StatusCode: http.StatusNotFound, StatusText: http.StatusText(http.StatusNotFound)}

// MethodNotAllowed 将结果定义为json格式
var MethodNotAllowed = JRouteReject{StatusCode: http.StatusMethodNotAllowed, StatusText: http.StatusText(http.StatusMethodNotAllowed)}

// JOK 工具函数：正常返回，无消息
var JOK = func(w http.ResponseWriter) {
	JComplete(w, http.StatusOK, code, message, nil)
}

// JResult 工具函数：正常返回，带消息
var JResult = func(w http.ResponseWriter, v interface{}) {
	JComplete(w, http.StatusOK, code, message, v)
}

// JReject 工具函数：非正常返回
var JReject = func(w http.ResponseWriter, status_code int, code int, message string) {
	JComplete(w, status_code, code, message, nil)
}
