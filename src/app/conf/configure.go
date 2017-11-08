package conf

import (
	"encoding/json"
	"io/ioutil"
	"libs/flag"
	"os"
)

type AppConfStruct struct {
	HttpHost string `json:"http.host"`
	HttpPort int    `json:"http.port"`

	WWWDir string `json:"www.dir"`
}

var AppConfInfo *AppConfStruct = nil

func init() {
	var path = flag.Parse("conf", "")

	if path == "" {
		panic("Need application configure file path")
	}

	var f, err_f = os.Open(path)
	if err_f != nil {
		panic(err_f)
	}

	defer func() { var _ = f.Close() }()

	var b, err_b = ioutil.ReadAll(f)
	if err_b != nil {
		panic(err_b)
	}

	var conf AppConfStruct
	var err_conf = json.Unmarshal(b, &conf)
	if err_conf != nil {
		panic(err_conf)
	}

	AppConfInfo = &conf
}
