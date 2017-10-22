package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var host = "127.0.0.1"
var port = 6379

func WithConn(host string, port int, fn func(v ...interface{})) {
	var c = fmt.Scanln(
}
