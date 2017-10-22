package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type JnRes struct {
	StatusCode string      `json:"status-code"`
	Message    string      `json:"message"`
	Error      string      `json:"error"`
	Body       interface{} `json:"body,omitempty"`
}

var default_code = 0
var default_message = ""
var default_error = ""

/*
 write json string to client

 status_code : http status code
 code : custom code
 message : http status message
 error : custom error message
 body : custom <struct body> message
*/
var JComplete = func(w http.ResponseWriter, status_code int, code int, message string, err string, body interface{}) {
	var msg = message
	if msg == "" {
		msg = http.StatusText(status_code)
	}

	var res = &JnRes{
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

var JOK = func(w http.ResponseWriter) {
	JComplete(w, http.StatusOK, default_code, default_message, default_error, nil)
}

var JResult = func(w http.ResponseWriter, v interface{}) {
	JComplete(w, http.StatusOK, default_code, default_message, default_error, v)
}

var JReject = func(w http.ResponseWriter, status_code int, code int, message string, err string) {
	JComplete(w, status_code, code, message, err, nil)
}
