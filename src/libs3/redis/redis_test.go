package redis

import (
	"testing"
)

func Test_WithConn(t *testing.T) {
	WithConn("127.0.0.1", 6379)(123)
}
