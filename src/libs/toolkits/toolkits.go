package toolkits

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

/*
golang interface -> json string without escape

params:
	v|interface

return:
	string
*/
func Encode(v interface{}) string {
	var buffer = bytes.NewBufferString("")

	var enc = json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	var _ = /*error ignore*/ enc.Encode(v)

	return buffer.String()
}

/*
json string -> golang type

params:
	js|string
	v|interface

return:
	error
*/
func Decode(js string, v interface{}) error {
	var dec = json.NewDecoder(bytes.NewBufferString(js))
	var err = dec.Decode(v)

	return err
}

/*
random number with seed(unix nano)
*/
func Random(n int) int {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var num = r.Intn(n)
	return num
}

/*
get the name of type
*/
func Type(v interface{}) (typ string) {
	switch vt := v.(type) {
	default:
		//val = fmt.Sprintf("%s", vt)
		typ = fmt.Sprintf("%T", vt)
	}
	return
}

func BinPath() string {
	var s, e = exec.LookPath(os.Args[0])
	if e != nil {
		panic(e)
	}

	var dir = filepath.Dir(s)
	return dir
}
