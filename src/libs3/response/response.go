package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonBody struct {
	StatusCode string      `json:"code"`
	Message    string      `json:"message"`
	Body       interface{} `json:"body,omitempty"`
}

var code = 0     /*default code*/
var message = "" /*default message*/

/*
 write json string to client

 status_code : http status code
 code : custom code
 message : client message or http status message or err msg
 body : custom <struct body> message
*/
var JComplete = func(w http.ResponseWriter, status_code int, code int, message string, body interface{}) {
	var msg = message
	if msg == "" {
		msg = http.StatusText(status_code)
	}

	var res = &JsonBody{
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

func DecoratorFunc(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer CatchError(w, r)
		//log.Info(r.Method, "  ", r.URL.Path, "  ", r.Header, "  ", r.Body)
		f(w, r)
		return
	}
}

var CatchError = func(w http.ResponseWriter, r *http.Request) {
	var v = recover()
	if v != nil {
		JReject(w, http.StatusInternalServerError, 0, fmt.Sprint(v))
	}
}

/*for rest, JSON*/
type JRouteReject struct {
	StatusCode int    /*http code*/
	StatusText string /*http message*/
}

func (rr *JRouteReject) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	JReject(w, rr.StatusCode, code, rr.StatusText)
}

var NotFound = JRouteReject{StatusCode: http.StatusNotFound, StatusText: http.StatusText(http.StatusNotFound)}
var MethodNotAllowed = JRouteReject{StatusCode: http.StatusMethodNotAllowed, StatusText: http.StatusText(http.StatusMethodNotAllowed)}

var JOK = func(w http.ResponseWriter) {
	JComplete(w, http.StatusOK, code, message, nil)
}

var JResult = func(w http.ResponseWriter, v interface{}) {
	JComplete(w, http.StatusOK, code, message, v)
}

var JReject = func(w http.ResponseWriter, status_code int, code int, message string) {
	JComplete(w, status_code, code, message, nil)
}
