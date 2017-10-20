package log

import (
	"fmt"
	"os"
	"testing"
)

type I interface {
	Say()
}

type Person struct {
	Name string
}

func (p *Person) Say() {
	fmt.Println(p.Name)
}

func (p *Person) Writer(pp []byte) (int, error) {
	fmt.Println(p.Name)
	return len(pp), nil
}

func Test_Log(t *testing.T) {
	var p = &Person{
		Name: "abc",
	}

	var is I = p
	is.Say()

	var s = os.Stdout
	fmt.Println(s.Fd(), s.Name())
	//s = os.Stderr
	//fmt.Println(s.Fd(), s.Name())
	//s = os.Stdin
	//fmt.Println(s.Fd(), s.Name())
	s, _ = os.OpenFile("app.log1", os.O_CREATE|os.O_RDWR, 0666)
	fmt.Println(s.Fd(), s.Name())
	s, _ = os.OpenFile("app.log", os.O_CREATE|os.O_RDWR, 0666)
	fmt.Println(s.Fd(), s.Name())
}
