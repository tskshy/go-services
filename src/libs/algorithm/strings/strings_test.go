package strings

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_BruteForce(t *testing.T) {
	var tt = "abc测试"
	var pp = "c测"
	var r = BruteForce(tt, pp)
	if !r {
		t.Error("alg error")
		return
	}

	tt = "abc测试"
	pp = "b测"
	r = BruteForce(tt, pp)
	if r {
		t.Error("alg error")
		return
	}
}

/*
 go test libs/algorithm/strings -v -test.run Test_MP
*/
func Test_MP(t *testing.T) {
	var ts = "ctcaatcacaatcat caatcat"
	//var ps = "caatcat"
	var ps = "adcadcad"
	t.Log(ts, ps)
	var r = MP(ts, ps, false)
	t.Log(r)
	if r != 0 {
		t.Error("alg error", r)
		return
	}

	r = MP(ts, ps, true)
	t.Log(r)
	if r != 0 {
		t.Error("alg error", r)
		return
	}
}

func Test_1(t *testing.T) {
	var m = make(map[string]string)
	m["hello"] = "1"
	var b, _ = json.Marshal(m)

	fmt.Println(string(b))
}
