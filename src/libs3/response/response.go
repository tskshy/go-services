package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonBody struct {
	StatusCode string      `json:"status-code"`
	Message    string      `json:"message"`
	Error      string      `json:"error,omitempty"`
	Body       interface{} `json:"body,omitempty"`
}

var default_code = 0     /*自定义code*/
var default_message = "" /*http message*/

var default_error = "" /*错误信息*/

/*
 write json string to client

 status_code : http status code
 code : custom code
 message : client message or http status message if message == ""
 error : custom error message
 body : custom <struct body> message
*/
var JComplete = func(w http.ResponseWriter, status_code int, code int, message string, err string, body interface{}) {
	var msg = message
	if msg == "" {
		msg = http.StatusText(status_code)
	}

	var res = &JsonBody{
		StatusCode: fmt.Sprintf("%d.%d", status_code, code),
		Message:    msg,
		Error:      err,
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

func DecoratorFunc(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	var ff = func(w http.ResponseWriter, r *http.Request) {
		defer CatchError(w, r)
		f(w, r)
		return
	}

	return ff
}

var CatchError = func(w http.ResponseWriter, r *http.Request) {
	var v = recover()
	if v != nil {
		JReject(w, http.StatusInternalServerError, 0, http.StatusText(http.StatusInternalServerError), fmt.Sprint(v))
	}
}

type RouteReject struct {
	StatusCode int    /*http code*/
	StatusText string /*http message*/
}

func (rr *RouteReject) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	JReject(w, rr.StatusCode, default_code, rr.StatusText, default_error)
}

var NotFound = RouteReject{StatusCode: http.StatusNotFound, StatusText: http.StatusText(http.StatusNotFound)}
var MethodNotAllowed = RouteReject{StatusCode: http.StatusMethodNotAllowed, StatusText: http.StatusText(http.StatusMethodNotAllowed)}

var JOK = func(w http.ResponseWriter) {
	JComplete(w, http.StatusOK, default_code, default_message, default_error, nil)
}

var JResult = func(w http.ResponseWriter, v interface{}) {
	JComplete(w, http.StatusOK, default_code, default_message, default_error, v)
}

var JReject = func(w http.ResponseWriter, status_code int, code int, message string, err string) {
	JComplete(w, status_code, code, message, err, nil)
}
