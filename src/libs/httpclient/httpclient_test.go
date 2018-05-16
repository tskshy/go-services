package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net"
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

func Test_TcpHttp(t *testing.T) {
	var conn, _ = net.Dial("tcp", "127.0.0.1:8080")

	var _, _ = conn.Write([]byte("POST /rongcloud/subscription/message?timestamp=2333268351272&nonce=1653208088&signature=266d33e77db1a6845a924452ca9d4be26db7e2c3&appKey=pgyu6atqyxixu HTTP/1.1\r\nHost: 127.0.0.1\r\n\r\nassss\r\n"))
	var b = true
	if b {
		conn.Close()
		return
	}

	var result = bytes.NewBuffer(nil)
	var buffer = make([]byte, 1024)
	var i = 0
	for {
		var n, errRead = conn.Read(buffer)
		fmt.Println(n, errRead)
		if errRead != nil {
			if errRead == io.EOF {
				break
			}
			return
		}
		result.Write(buffer[0:n])
		i++
		if i == 1 {
			conn.Close()
			return
		}
	}

	conn.Close()

	fmt.Println(string(result.Bytes()))
}
