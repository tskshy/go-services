package flag

import (
	"fmt"
	"testing"
)

func Test_Parse(t *testing.T) {
	fmt.Println(Parse("test.v=true", "asd"))
}
