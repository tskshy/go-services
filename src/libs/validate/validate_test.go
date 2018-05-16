package validate

import (
	"fmt"
	"testing"
)

type User struct {
	ID   string `validate:"^[a-z]{1,2}$" validate-message:"账户非法"`
	Name string `validate:""`
	Age  int    `validate:"^([0]|[1][0-9]{0,2})$"`
}

func Test_Tag(t *testing.T) {
	var user = &User{
		ID:  "abcqwe",
		Age: 20,
	}
	fmt.Println(Validate(*user))
	//var b, _ = json.Marshal(user)
	//fmt.Println(string(b))
}

func Benchmark_v(b *testing.B) {
	var user = &User{
		ID:  "abcqwe",
		Age: 20,
	}

	for i := 0; i < b.N; i++ {
		Validate(*user)
	}
}

func Benchmark_v1(b *testing.B) {
	var user = &User{
		ID:  "abcqwe",
		Age: 20,
	}

	for i := 0; i < b.N; i++ {
		Validate1(*user)
	}
}
