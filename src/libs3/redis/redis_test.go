package redis

import (
	"testing"
)

/*usage*/
func Test_Redis(t *testing.T) {
	var c = Pool.Get()
	if c.Err() != nil {
		t.Error(c.Err())
		return
	}

	var e = c.Close()
	if e != nil {
		t.Error(e)
		return
	}
	return
}

func Test_WithConn(t *testing.T) {

}
