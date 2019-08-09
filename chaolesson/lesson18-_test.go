package chaolesson

import (
	"fmt"
	"github.com/spotmaxtech/gokit"
	"testing"
)

type L18Programmer interface {
	WriteHelloWorld() Code
}

type L18JavaProgrammer struct {
}

func (p *L18JavaProgrammer) WriteHelloWorld() Code {
	return "java hello world"
}

type L18GoProgrammer struct {
}

func (p *L18GoProgrammer) WriteHelloWorld() Code {
	return "go hello world"
}

func HelloWorld(p L18Programmer) {
	fmt.Println(p.WriteHelloWorld())
}

func TestL18_programmer(t *testing.T) {
	goPro := new(L18GoProgrammer)
	javaPro := new (L18JavaProgrammer)
	HelloWorld(goPro)
	HelloWorld(javaPro)
}

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("int", i)
	} else {
		fmt.Println("ops", i)
	}

	switch p.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func TestSomething(t *testing.T) {
	DoSomething(5)
	DoSomething("hello")
}

func TestRangeFault(t *testing.T) {
	l1 := []int{1,2,3,4,5}
	var l2 []*int
	for _, i := range l1 {
		l2 = append(l2, &i)
	}

	t.Log(gokit.Prettify(l2))
}