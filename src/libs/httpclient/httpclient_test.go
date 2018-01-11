package httpclient

import (
	"testing"
)

/*
 go test libs/httpclient -v -test.run Test_HttpClient
*/
func Test_HttpClient(t *testing.T) {
	var request = &HttpRequest{
		Method: "POST",
		Url:    "http://127.0.0.1:11111",
	}

	var body, err = request.Do()

	if err != nil {
		t.Error(err)
	} else {
		t.Log(body)
	}
}
