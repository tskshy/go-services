package strings

import (
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
	var ts = "ctcaatcacaatcat"
	//var ps = "caatcat"
	//var ps = "abaabbabaab"
	var ps = "abcdabd"
	t.Log(ts, ps)
	var r = MP1(ts, ps)
	if !r {
		t.Error("alg error", r)
		return
	}
}
