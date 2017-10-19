package httpclient

import (
	BYTES "bytes"
	JSON "encoding/json"
	ERRORS "errors"
	FMT "fmt"
	IO "io"
	IOUTIL "io/ioutil"
	MULTIPART "mime/multipart"
	HTTP "net/http"
	URL "net/url"
	OS "os"
	FILEPATH "path/filepath"
	STRINGS "strings"
	TIME "time"
)

type HttpRequest struct {
	Method  string            //请求方法
	Url     string            //请求URL
	Params  map[string]string //URL上的参数
	Headers map[string]string //请求头
	Body    string            //请求实体
	Timeout int               //超时设置(秒)
}

/*
 * HTTP CLIENT
 */
func request(hr *HttpRequest) (*HTTP.Response, error) {
	var info = *hr

	u, err := URL.Parse(info.Url)
	if err != nil {
		return nil, err
	}

	var pms = u.Query()
	if nil != info.Params && len(info.Params) > 0 {
		for k, v := range info.Params {
			pms.Set(k, v)
		}
	}

	var url_buffer BYTES.Buffer
	url_buffer.WriteString(u.Scheme)
	url_buffer.WriteString("://")
	url_buffer.WriteString(u.Host)
	url_buffer.WriteString(u.Path)

	if len(pms) > 0 {
		url_buffer.WriteString("?")
		url_buffer.WriteString(pms.Encode())
	}

	var body IO.Reader = nil

	if info.Body != "" {
		if info.Headers != nil &&
			len(info.Headers) > 0 &&
			info.Headers["Content-Type"] == "multipart/form-data" {

			/*解析info.Body，要求必须是json格式|map[string]string*/
			//var tmp = `{"k": "v", "k1": "@/root/devel/golang/test/active.sh"}`
			var m = make(map[string]string)
			var err = JSON.Unmarshal([]byte(info.Body), &m)
			if err != nil {
				return nil, err
			}
			var multipart_body = &BYTES.Buffer{}
			var writer = MULTIPART.NewWriter(multipart_body)

			err = func() error {
				defer func() { var _ = writer.Close() }()

				for k, v := range m {
					if STRINGS.HasPrefix(v, "@") {
						var file_path = v[1:]

						file, err := OS.Open(file_path)
						if err != nil {
							return err
						}
						defer func() { var _ = file.Close() }()

						part, err := writer.CreateFormFile(k, FILEPATH.Base(file_path))
						if err != nil {
							return err
						}

						_, err = IO.Copy(part, file)
						if err != nil {
							return err
						}
					} else {
						err := writer.WriteField(k, v)
						if err != nil {
							return err
						}
					}
				}
				return nil
			}()
			if err != nil {
				return nil, err
			}

			info.Headers["Content-Type"] = writer.FormDataContentType()
			body = STRINGS.NewReader(multipart_body.String())
		} else {
			body = STRINGS.NewReader(info.Body)
		}
	}

	request, err := HTTP.NewRequest(STRINGS.ToUpper(info.Method), url_buffer.String(), body)
	if err != nil {
		return nil, err
	}

	if info.Headers != nil && len(info.Headers) > 0 {
		for k, v := range info.Headers {
			request.Header.Set(k, v)
		}
	}

	var client = &HTTP.Client{
		Timeout: TIME.Duration(TIME.Duration(info.Timeout) * TIME.Second),
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
 * params: HttpRequest
 * http请求信息封装，仅接受字符串，文件下载函数暂无
 * return: (string, error)
 */
func (hr *HttpRequest) Do() (string, error) {
	response, err := request(hr)
	if err != nil {
		return "", err
	}

	defer func() { var _ = response.Body.Close() }()

	resbody, err := IOUTIL.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return string(resbody), ERRORS.New(`http status code: ` + FMT.Sprintf("%d", response.StatusCode))
	}

	return string(resbody), nil
}
